package dao

import (
	"context"
	"guoshao-fm-crawler/internal/model/entity"
	"strconv"
	"testing"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func TestInsertFeedChannel(t *testing.T) {
	var (
		testModel entity.FeedChannel
	)

	g.DB().SetDryRun(true)
	testModel = entity.FeedChannel{
		Title:       "test title",
		ChannelDesc: "test channelDesc",
		ImageUrl:    "https://www.test.com/test.png",
		Link:        "https://www.test.com/test",
		FeedLink:    "https://www.test.com/test.xml",
		Copyright:   "test copyright",
		Language:    "zh-cn",
		Author:      "test",
		OwnerName:   "test",
		OwnerEmail:  "test",
		FeedType:    "test",
		Categories:  "test1,test2",
	}
	testModel.Id = strconv.FormatUint(ghash.RS64([]byte(testModel.FeedLink+testModel.Title)), 32)
	type args struct {
		ctx   context.Context
		model entity.FeedChannel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Insert feed channel ",
			args: args{
				ctx:   gctx.New(),
				model: testModel,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertFeedChannelIfNotExist(tt.args.ctx, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("InsertFeedChannel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFeedChannelList(t *testing.T) {
	type args struct {
		ctx    context.Context
		offset int
		limit int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get channel list from db",
			args: args{
				ctx:    gctx.New(),
				offset: 0,
				limit: 10,
			},
		},
	}
	g.DB().SetDryRun(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := GetFeedChannelList(tt.args.ctx, tt.args.offset, tt.args.limit)
			if len(list) < 10 {
				t.Fatal("The feed channel list size is less than offset 10")
			}
			t.Log("The result feed channel list size is : ", len(list))
		})
	}
}

func TestGetAllFeedChannelCount(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name      string
		args      args
	}{
		{
			name: "Get channel total count",
			args: args{
				ctx: gctx.New(),
			},
		},
	}
	g.DB().SetDryRun(true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count, err := GetAllFeedChannelCount(tt.args.ctx)
			if err != nil {
				t.Fatal("Get channel total count failed : ", err)
			}
			t.Log("The channel total count is ", count)
		})
	}
}
