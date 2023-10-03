package worker

import (
	"context"
	"porkast-crawler/internal/consts"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

func Test_getPodbeanCategoryHtml(t *testing.T) {
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get podbean all category html string",
			args: args{
				ctx: gctx.New(),
				url: consts.PODBEAN_ALL_CATEGORY_URL,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respStr := getPodbeanCategoryHtml(tt.args.ctx, tt.args.url)
			if respStr == "" {
				t.Fatal("The podbean all category html response is empty")
			}
		})
	}
}

func Test_getPodbeanCategoryUrlsFromHtml(t *testing.T) {
	type args struct {
		ctx             context.Context
		categoryHtmlStr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Parse podbean all category url list",
			args: args{
				ctx:             gctx.New(),
				categoryHtmlStr: gfile.GetContents("./testdata/podben_all.html"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			urlList := getPodbeanCategoryUrlsFromHtml(tt.args.ctx, tt.args.categoryHtmlStr)
			if len(urlList) == 0 {
				t.Fatal("Parse podbean category url list from html failed, the url list is empty")
			}
		})
	}
}

func Test_getPodbeanCategoryPopularShowHtml(t *testing.T) {
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Parse podbean category popular show list",
			args: args{
				ctx: gctx.New(),
				url: "https://www.podbean.com/site/category/getPopularList?page=2&category=Religion+%26+Spirituality",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				respStr string
			)

			respStr = getPodbeanCategoryPopularShowHtml(tt.args.ctx, tt.args.url)
			if respStr == "" {
				t.Fatal("The response is empty")
			}
		})
	}
}

func Test_getPodbeanCategoryPopularShowRSSList(t *testing.T) {
	type args struct {
		ctx     context.Context
		htmlStr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get podbean RSS link",
			args: args{
				ctx:     gctx.New(),
				htmlStr: gfile.GetContents("./testdata/podbean_category_popular.html"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				rssLinkList []string
			)
			rssLinkList = getPodbeanCategoryPopularShowRSSList(tt.args.ctx, tt.args.htmlStr)
			if len(rssLinkList) == 0 {
				t.Fatal("The podbean RSS link list is empty")
			}
		})
	}
}
