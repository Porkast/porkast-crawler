package worker

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

func TestParseFirstoryCategoryItemList(t *testing.T) {
	type args struct {
		categoryId string
		skip       int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseFirstoryCategoryItemList(tt.args.categoryId, tt.args.skip)
		})
	}
}

func Test_parseFirstoryAllCategoryList(t *testing.T) {
	type args struct {
		jsonStr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Parse all firstory category list",
			args: args{
				jsonStr: gfile.GetContents("./testdata/firsotry_category_list.json"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var categoryIdList []string
			categoryIdList = parseFirstoryAllCategoryList(tt.args.jsonStr)
			if len(categoryIdList) == 0 {
				t.Fatal("Parse firstory category json list failed")
			}
		})
	}
}

func Test_getFirstoryAllCategoryList(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get firstory all category by graphql",
			args: args{
				ctx: gctx.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var resp string
			resp = getFirstoryAllCategoryList(tt.args.ctx)
			if resp == "" {
				t.Fatal("Get firstory all category by graphql failed")
			}
		})
	}
}

func Test_getFirstoryCategoryShowList(t *testing.T) {
	type args struct {
		ctx        context.Context
		categoryId string
		skip       int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get firstory category show list by post graphql query string",
			args: args{
				ctx:        gctx.New(),
				categoryId: "ck0zfa8rdcd370786j56unogu",
				skip:       40,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				resp     string
				respJson *gjson.Json
			)
			resp = getFirstoryCategoryShowsJsonStr(tt.args.ctx, tt.args.categoryId, tt.args.skip)
			if resp == "" {
				t.Fatal("The response string is empty")
			}

			respJson = gjson.New(resp)
			if respJson == nil || respJson.IsNil() {
				t.Fatal("The response json string is not valid")
			}

		})
	}
}

func Test_parseFirstoryShowList(t *testing.T) {
	type args struct {
		jsonStr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Parse firstory show id list",
			args: args{
				jsonStr: gfile.GetContents("./testdata/firstory_show_list.json"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				showIdList []string
			)
			showIdList = parseFirstoryShowList(tt.args.jsonStr)
			if len(showIdList) == 0 {
				t.Fatal("The show id list is empty")
			}
		})
	}
}

func Test_getFirstoryCategoryShowInfoJsonStr(t *testing.T) {
	type args struct {
		ctx        context.Context
		categoryId string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get firstory show item json string without external RSS",
			args: args{
				ctx:        gctx.New(),
				categoryId: "cjfl8pzko3fwb0192dfgrn1so",
			},
		},
		{
			name: "Get firstory show item json string with external RSS",
			args: args{
				ctx:        gctx.New(),
				categoryId: "ipodcast",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				respStr  string
				respJson *gjson.Json
			)
			respStr = getFirstoryCategoryShowInfoJsonStr(tt.args.ctx, tt.args.categoryId)
			if respStr == "" {
				t.Fatal("The response string is empty")
			}
			respJson = gjson.New(respStr)
			if respJson == nil || respJson.IsNil() {
				t.Fatal("The response json is empty")
			}
		})
	}
}

func Test_getFirsotryShowRSSLink(t *testing.T) {
	type args struct {
		ctx     context.Context
		jsonStr string
	}
	tests := []struct {
		name           string
		args           args
		isExternalLink bool
	}{
		{
			name: "Parse internal show RSS",
			args: args{
				ctx:     gctx.New(),
				jsonStr: gfile.GetContents("./testdata/firsotry_category_list.json"),
			},
			isExternalLink: false,
		},
		{
			name: "Parse external show RSS",
			args: args{
				ctx:     gctx.New(),
				jsonStr: gfile.GetContents("./testdata/firstory_show_resp_external_rss.json"),
			},
			isExternalLink: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				rssLink string
			)
			rssLink = getFirsotryShowRSSLink(tt.args.ctx, tt.args.jsonStr)
			if tt.isExternalLink && gstr.HasPrefix(rssLink, "https://open.firstory.me/rss/user") {
				t.Fatal("Parse internal RSS link failed")
			}
		})
	}
}
