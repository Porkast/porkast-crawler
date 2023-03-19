package worker

import (
	"guoshao-fm-crawler/internal/model/entity"
	"strconv"
	"testing"

	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/mmcdole/gofeed"
)

func Test_feedChannelToModel(t *testing.T) {
	var (
		err     error
		feedStr string
		feedID  string
		feed    *gofeed.Feed
		fp      *gofeed.Parser
		model   entity.FeedChannel
	)

	feedStr = gfile.GetContents("./testdata/ximalaya_rss_resp.xml")
	fp = gofeed.NewParser()
	feed, err = fp.ParseString(feedStr)
	if err != nil {
		t.Fatal("Parse feed string failed : ", err)
	}

	feedID = strconv.FormatUint(ghash.RS64([]byte(feed.FeedLink+feed.Title)), 32)
	model = feedChannelToModel(feedID, *feed)
	if model.Id != feedID || model.Title != "李峙的不老歌" || model.Author != "李峙" || model.Copyright != "李峙电台 @喜马拉雅FM" || model.FeedLink != "https://www.ximalaya.com/album/236268.xml" || model.Link != "https://www.ximalaya.com" || model.OwnerEmail != "radiolizhi@163.com" {
		t.Fatal("The channel info is not correct")
	}

}
