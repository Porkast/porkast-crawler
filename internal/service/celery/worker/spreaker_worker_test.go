package worker

import (
	"context"
	"os"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

func Test_getSpreakerCategory(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get Spreaker category page",
			args: args{
				ctx: gctx.New(),
			},
		},
	}
	if os.Getenv("env") != "dev" {
		return
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				respStr string
			)
			respStr = getSpreakerCategory(tt.args.ctx)
			if respStr == "" {
				t.Fatal("Get Spreaker category page failed")
			}
		})
	}
}

func Test_parseSpreakerCategoryItemList(t *testing.T) {
	type args struct {
		ctx             context.Context
		categoryHtmlStr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Parse Spreaker category item list",
			args: args{
				ctx:             gctx.New(),
				categoryHtmlStr: gfile.GetContents("./testdata/spreaker_category.html"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				categoryList []string
			)
			categoryList = parseSpreakerCategoryItemList(tt.args.ctx, tt.args.categoryHtmlStr)
			if len(categoryList) == 0 {
				t.Fatal("The sprealer category list is empty")
			}

			for _, categoryStr := range categoryList {
				if categoryStr == "" {
					t.Fatal("The category item is empty")
				}
			}
		})
	}
}

func Test_getSpreakerCategoryNextUrl(t *testing.T) {
	type args struct {
		respStr string
	}
	tests := []struct {
		name        string
		args        args
		wantNextUrl string
	}{
		{
			name: "Get spreaker next url",
			args: args{
				respStr: gfile.GetContents("./testdata/spreaker_category_arts.html"),
			},
			wantNextUrl: "https://api.spreaker.com/v2/explore/categories/arts/items?l=zh&export=show_author%2Cshow_profile&last_id=3497433&limit=50",
		},
		{

			name: "Get spreaker next url from json",
			args: args{
				respStr: gfile.GetContents("./testdata/spreaker_more_category_resp.json"),
			},
			wantNextUrl: "https://api.spreaker.com/v2/explore/categories/92/items?l=zh&last_id=1703669&limit=50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNextUrl := getSpreakerCategoryNextUrl(tt.args.respStr); gotNextUrl != tt.wantNextUrl {
				t.Errorf("getSpreakerCategoryNextUrl() = %v, want %v", gotNextUrl, tt.wantNextUrl)
			}
		})
	}
}

func Test_getSpreakerShowLinks(t *testing.T) {
	type args struct {
		respStr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get spreaker show link list",
			args: args{
				respStr: gfile.GetContents("./testdata/spreaker_category_arts.html"),
			},
		},
		{
			name: "Get spreaker show link list from json",
			args: args{
				respStr: gfile.GetContents("./testdata/spreaker_more_category_resp.json"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linkList := getSpreakerShowLinks(tt.args.respStr)
			if len(linkList) == 0 {
				t.Fatal("The spreaker show link list is empty")
			}
		})
	}
}

func Test_getSpreakerShowRSSLink(t *testing.T) {
	type args struct {
		resp string
	}
	tests := []struct {
		name        string
		args        args
		wantRssLink string
	}{
		{
			name: "Get RSS link from show response",
			args: args{
				resp: gfile.GetContents("./testdata/spreaker_show_resp.html"),
			},
			wantRssLink: "https://www.spreaker.com/show/3497433/episodes/feed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRssLink := getSpreakerShowRSSLink(tt.args.resp); gotRssLink != tt.wantRssLink {
				t.Errorf("getSpreakerShowRSSLink() = %v, want %v", gotRssLink, tt.wantRssLink)
			}
		})
	}
}
