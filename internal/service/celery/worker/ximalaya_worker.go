package worker

import (
	"fmt"
	"guoshao-fm-crawler/internal/service/celery/jobs"
	"guoshao-fm-crawler/internal/service/network"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func ParseXiMaLaYaPodcast(url string) {
	var (
		ctx        = gctx.New()
		podcastUrl string
		respStr    string
	)
	time.Sleep(time.Second * 3)
	podcastUrl = url + ".xml"
    g.Log().Line().Debug(ctx, "Start parse ximalaya podcast with url : ", podcastUrl)
	respStr = network.TryGetRSSContent(ctx, podcastUrl)
	if isStringRSSXml(respStr) {
		//The ximalaya album is RSS
		storeFeed(ctx, respStr)
	}
}

func ParseXiMaLaYaEntry(url string) {
	var (
		ctx          = gctx.New()
		respStr      string
		respJson     *gjson.Json
		albumUrlList []string
	)

    g.Log().Line().Debug(ctx, "Start parse ximalaya entry")
	respStr = network.GetContent(ctx, url)
	respJson = gjson.New(respStr)
	if respJson == nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Parse ximalaya albums response json failed.\nUrl is %s\nResponse String is %s", url, respStr))
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
