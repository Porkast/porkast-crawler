package worker

import (
	"testing"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

func Test_getLizhiCategory(t *testing.T) {
	var (
		// ctx             = gctx.New()
		lizhiHomeHtml   string
		categoryUrlList []string
	)

	lizhiHomeHtml = gfile.GetContents("./testdata/lizhi_home_page.html")
	categoryUrlList = getLizhiCategory(lizhiHomeHtml)
	if len(categoryUrlList) == 0 {
		t.Fatal("Get Lizhi categoryUrlList failed")
	}

	for _, url := range categoryUrlList {
		if !gstr.Contains(url, "www.lizhi.fm/label") {
			t.Fatal("Get Lizhi categoryUrlList failed, the url not contain lizhi domain")
		}

		if gstr.HasPrefix(url, "//") {
			t.Fatal("Get Lizhi categoryUrlList failed, the url has prefix //")
		}

		if !gstr.HasSuffix(url, ".html") {
			t.Fatal("Get Lizhi categoryUrlList failed, the url has no suffix .html")
		}
	}

}

func Test_getLizhiPodcastIdList(t *testing.T) {
	var (
		htmlStr       string
		podcastIdList []string
	)

	htmlStr = gfile.GetContents("./testdata/lizhi_category_page.html")
	podcastIdList = getLizhiPodcastIdList(htmlStr)
	if len(podcastIdList) == 0 {
		t.Fatal("Parse podcast id list from category page failed")
	}

}

func Test_getLizhiCurrentCategoryPageRadioCount(t *testing.T) {
	var (
		htmlStr        string
		endPageHtmlStr string
		count          int
	)

	htmlStr = gfile.GetContents("./testdata/lizhi_category_page.html")
	count = getLizhiCurrentCategoryPageRadioCount(htmlStr)
	if count != 24 {
		t.Fatal("Get lizhi category total radio count failed, should be 24, but it is ", count)
	}

	endPageHtmlStr = gfile.GetContents("./testdata/lizhi_category_end_page.html")
	count = getLizhiCurrentCategoryPageRadioCount(endPageHtmlStr)
	if count != 2 {
		t.Fatal("Get lizhi category total radio count failed, should be 2, but it is ", count)
	}
}

func Test_getLizhiCategoryCurrentPageNum(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantNum int
	}{
		{
			name: "Get page number from url",
			args: args{
				url: "https://www.lizhi.fm/label/24229874933174064/4.html",
			},
			wantNum: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNum := getLizhiCategoryCurrentPageNum(tt.args.url); gotNum != tt.wantNum {
				t.Errorf("getLizhiCategoryCurrentPageNum() = %v, want %v", gotNum, tt.wantNum)
			}
		})
	}
}

func Test_genNextLizhiCategoryPageUrl(t *testing.T) {
	type args struct {
		url         string
		currentPage int
	}
	tests := []struct {
		name        string
		args        args
		wantNextUrl string
	}{
		{
			name: "Get next page count",
			args: args{
				url:         "https://www.lizhi.fm/label/24229874933174064/1.html",
				currentPage: 1,
			},
			wantNextUrl: "https://www.lizhi.fm/label/24229874933174064/2.html",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNextUrl := genNextLizhiCategoryPageUrl(tt.args.url, tt.args.currentPage); gotNextUrl != tt.wantNextUrl {
				t.Errorf("genNextLizhiCategoryPageUrl() = %v, want %v", gotNextUrl, tt.wantNextUrl)
			}
		})
	}
}

func Test_genLizhiRSSUrl(t *testing.T) {
	type args struct {
		podcastId string
	}
	tests := []struct {
		name    string
		args    args
		wantUrl string
	}{
		{
			name: "Generate Lizhi FM RSS url",
			args: args{
				podcastId: "1066826",
			},
			wantUrl: "https://rss.lizhi.fm/rss/1066826.xml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUrl := genLizhiRSSUrl(tt.args.podcastId); gotUrl != tt.wantUrl {
				t.Errorf("genLizhiRSSUrl() = %v, want %v", gotUrl, tt.wantUrl)
			}
		})
	}
}
