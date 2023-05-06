package worker

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

func Test_parseAllApplePodcastCategoryUrlList(t *testing.T) {
	type args struct {
		ctx     context.Context
		htmlStr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "parse all category url list",
			args: args{
				ctx:     gctx.New(),
				htmlStr: gfile.GetContents("./testdata/apple_podcast_entry_resp.html"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urlList := parseAllApplePodcastCategoryUrlList(tt.args.ctx, tt.args.htmlStr)
			if len(urlList) == 0 {
				t.Fatal("parse all apple podcast category url list failed")
			}
		})
	}
}

func Test_parseApplePodcastCategoryItemUrls(t *testing.T) {
	type args struct {
		ctx     context.Context
		htmlStr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "parse category item url list",
			args: args{
				ctx:     gctx.New(),
				htmlStr: gfile.GetContents("./testdata/apple_podcast_category_item.html"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itemUrlList := parseApplePodcastCategoryItemUrls(tt.args.ctx, tt.args.htmlStr)
			if len(itemUrlList) == 0 {
				t.Fatal("parse apple podcast category item url list failed")
			}
		})
	}
}

func Test_parseApplePodcastItemId(t *testing.T) {
	type args struct {
		ctx     context.Context
		itemUrl string
	}
	tests := []struct {
		name       string
		args       args
		wantItemId string
	}{
		{
			name: "parse apple podcast item id by item url",
			args: args{
				ctx:     gctx.New(),
				itemUrl: "https://podcasts.apple.com/cn/podcast/%E5%B0%8F%E8%AF%B4-%E6%83%85%E4%B9%A6-%E4%BD%9C%E8%80%85-%E5%B2%A9%E4%BA%95%E4%BF%8A%E4%BA%8C-%E6%9C%97%E8%AF%BB-%E5%B0%8F%E5%B2%9B/id1232948544",
			},
			wantItemId: "1232948544",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotItemId := parseApplePodcastItemId(tt.args.ctx, tt.args.itemUrl); gotItemId != tt.wantItemId {
				t.Errorf("parseApplePodcastItemId() = %v, want %v", gotItemId, tt.wantItemId)
			}
		})
	}
}

func Test_getApplePodcastItemRSSByLookupAPI(t *testing.T) {
	type args struct {
		ctx    context.Context
		itemId string
	}
	tests := []struct {
		name    string
		args    args
	}{
		{
			name: "get rss content from apple podcast itune api",
			args: args{
				ctx: gctx.New(),
				itemId: "1232948544",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rssContent := getApplePodcastItemRSSByLookupAPI(tt.args.ctx, tt.args.itemId)
			if rssContent == "" {
				t.Fatal("get rss content from apple podcast api failed")
			}
		})
	}
}
