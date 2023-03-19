package worker

import (
	"fmt"
	"guoshao-fm-crawler/internal/model/entity"
	"guoshao-fm-crawler/internal/service/celery/jobs"
	"guoshao-fm-crawler/internal/service/internal/dao"
	"guoshao-fm-crawler/internal/service/network"
	"guoshao-fm-crawler/utility"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/mmcdole/gofeed"
)

func ParseXiMaLaYaPodcast(url string) {
	var (
		ctx        = gctx.New()
		podcastUrl string
		respStr    string
		feed       *gofeed.Feed
	)
	time.Sleep(time.Second * 3)
	podcastUrl = url + ".xml"
	respStr = network.GetContent(ctx, podcastUrl)
	if isStringXml(respStr) {
		//The ximalaya album is RSS
		feed = utility.ParseFeed(ctx, respStr)
		if feed != nil {
			g.Log().Info(ctx, "Get feed %s ")
			var (
				feedChannelMode entity.FeedChannel
				feedItemList    []entity.FeedItem
				feedID          string
			)
			feedID = strconv.FormatUint(ghash.RS64([]byte(feed.FeedLink+feed.Title)), 32)
			feedChannelMode = feedChannelToModel(feedID, *feed)
			for _, item := range feed.Items {
				var (
					feedItem entity.FeedItem
				)
				feedItem = feedItemToModel(feedID, *item)
				feedItemList = append(feedItemList, feedItem)
			}
			dao.InsertFeedChannelIfNotExist(ctx, feedChannelMode)
			for _, item := range feedItemList {
				dao.InsertFeedItemIfNotExist(ctx, item)
			}
		}
	}
}


func ParseXiMaLaYaEntry(url string) {
	var (
		ctx          = gctx.New()
		respStr      string
		respJson     *gjson.Json
		albumUrlList []string
	)

	respStr = network.GetContent(ctx, url)
	respJson = gjson.New(respStr)
	if respJson == nil {
		g.Log().Error(ctx, fmt.Sprintf("Parse ximalaya albums response json failed.\nUrl is %s\nResponse String is %s", url, respStr))
	}

	albumUrlList = getXimalayaAlbumUrlList(*respJson)
	for _, albumUrl := range albumUrlList {
		jobs.AssignXimalayaPodcastJob(ctx, albumUrl)
	}
}

func getXimalayaAlbumUrlList(data gjson.Json) (albumUrlList []string) {
	var (
		ximalayaBaseUrl = "https://www.ximalaya.com"
		albumJsons      []*gjson.Json
	)

	albumJsons = data.GetJsons("data.albums")
	for _, albumJson := range albumJsons {
		var (
			albumUrl string
		)

		albumUrl = albumJson.Get("albumUrl").String()
		albumUrlList = append(albumUrlList, ximalayaBaseUrl+albumUrl)
	}

	return
}
