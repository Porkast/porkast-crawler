package worker

import (
	"context"
	"porkast-crawler/internal/model/entity"
	"strconv"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/mmcdole/gofeed"
)

func Test_isStringRSSXml(t *testing.T) {
	type args struct {
		respStr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Response string is RSS case",
			args: args{
				respStr: gfile.GetContents("./testdata/ximalaya_rss_resp.xml"),
			},
			want: true,
		},
		{
			name: "Response html string is not RSS case",
			args: args{
				respStr: gfile.GetContents("./testdata/lizhi_404.html"),
			},
			want: false,
		},
		{
			name: "Response string is not RSS case",
			args: args{
				respStr: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStringRSSXml(tt.args.respStr); got != tt.want {
				t.Errorf("isXimalayaRespXml() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func Test_setChannelLastUpdateRecord(t *testing.T) {
	type args struct {
		ctx       context.Context
		channelId string
		funName   string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "set channel last update time",
			args: args{
				ctx:       gctx.New(),
				channelId: "10snekvib3e4i",
				funName:   "Test_setChannelLastUpdateRecord",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setChannelLastUpdateRecord(tt.args.ctx, tt.args.channelId, tt.args.funName)
		})
	}
}

func Test_storeFeed(t *testing.T) {
	type args struct {
		ctx      context.Context
		respStr  string
		feedLink string
		funName  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "store feed",
			args: args{
				ctx:      gctx.New(),
				respStr:  gfile.GetContents("./testdata/apple_podcast_npe_rss.xml"),
				feedLink: "http://www.mysteryshows.com/",
				funName:  "Test_storeFeed",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storeFeed(tt.args.ctx, tt.args.respStr, tt.args.feedLink, tt.args.funName)
		})
	}
}
