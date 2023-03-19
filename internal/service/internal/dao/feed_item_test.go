package dao

import (
	"context"
	"guoshao-fm-crawler/internal/model/entity"
	"strconv"
	"testing"

	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

func TestInsertFeedItemIfNotExist(t *testing.T) {
	var (
		testModel entity.FeedItem
	)

	testModel = entity.FeedItem{
		Id:              "85url7dhbiu9o",
		ChannelId:       "cako6c4r1hdjm",
		Title:           "【李峙的乐未央】天空，是头顶的大海",
		Link:            "https://www.ximalaya.com//1716986/sound/618172108",
		PubDate:         gtime.New(),
		Author:          "",
		InputDate:       gtime.New(),
		ImageUrl:        "https://fdfs.xmcdn.com/storages/daf7-audiofreehighqps/02/89/GMCoOSQH5kPzAAiKJgICe9WI.jpeg",
		EnclosureUrl:    "https://jt.ximalaya.com//GKwRIRwH5kRtAL8qwwICfBRc.m4a?channel=rss&album_id=236268&track_id=618172108&uid=1716986&jt=https://audio.xmcdn.com/storages/ba06-audiofreehighqps/8F/E9/GKwRIRwH5kRtAL8qwwICfBRc.m4a",
		EnclosureType:   "audio/x-m4a",
		EnclosureLength: "61898997",
		Duration:        "25:47",
		Episode:         "1",
		Explicit:        "",
		Season:          "1",
	}
	testModel.Id = strconv.FormatUint(ghash.RS64([]byte(testModel.Link+testModel.Title)), 32)

	g.DB().SetDryRun(true)
	type args struct {
		ctx   context.Context
		model entity.FeedItem
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Insert feed item",
			args: args{
				ctx:   gctx.New(),
				model: testModel,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertFeedItemIfNotExist(tt.args.ctx, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("InsertFeedItemIfNotExist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
