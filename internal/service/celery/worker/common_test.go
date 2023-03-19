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

func Test_feedItemToModel(t *testing.T) {
	var (
		err     error
		feedStr string
		feedID  string
		feed    *gofeed.Feed
		fp      *gofeed.Parser
		model   entity.FeedItem
	)

	feedStr = gfile.GetContents("./testdata/ximalaya_rss_resp.xml")
	fp = gofeed.NewParser()
	feed, err = fp.ParseString(feedStr)
	if err != nil {
		t.Fatal("Parse feed string failed : ", err)
	}

	feedID = strconv.FormatUint(ghash.RS64([]byte(feed.FeedLink+feed.Title)), 32)
	model = feedItemToModel(feedID, *feed.Items[0])
	if model.Title != "【李峙的乐未央】天空，是头顶的大海" {
		t.Fatal("The item title is not corret")
	}

	if model.EnclosureUrl != "https://jt.ximalaya.com//GKwRIRwH5kRtAL8qwwICfBRc.m4a?channel=rss&album_id=236268&track_id=618172108&uid=1716986&jt=https://audio.xmcdn.com/storages/ba06-audiofreehighqps/8F/E9/GKwRIRwH5kRtAL8qwwICfBRc.m4a" {
		t.Fatal("The item Enclosure url is not corret")
	}

	if model.EnclosureLength != "61898997" {
		t.Fatal("The item Enclosure lenght is not corret")
	}

	if model.EnclosureType != "audio/x-m4a" {
		t.Fatal("The item Enclosure type is not corret")
	}

	if model.Duration != "25:47" {
		t.Fatal("The item Duration is not corret")
	}

	if model.Link != "https://www.ximalaya.com//1716986/sound/618172108" {
		t.Fatal("The item Link is not corret")
	}
}
