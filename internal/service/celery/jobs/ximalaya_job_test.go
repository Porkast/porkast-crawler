package jobs

import (
	"testing"

	"github.com/gogf/gf/v2/text/gstr"
)

func Test_getEntryUrlList(t *testing.T) {
	var (
		urlList []string
	)

	urlList = getXimalayaEntryUrlList()
	if len(urlList) == 0 {
		t.Fatal("ximalaya entru url list is empty")
	}

	for _, url := range urlList {
		var (
			targetUrl string
		)
		targetUrl = formatXimalayaUrl(url, 1)
		if !gstr.Contains(targetUrl, "pageNum=1") {
			t.Fatal("Format url with page number failed")
		}
	}
}

func Test_formatXimalayaUrl(t *testing.T) {
	type args struct {
		url  string
		page int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "format ximalaya url",
			args: args{
				url: "https://www.ximalaya.com/revision/category/v2/albums?pageNum=%d&pageSize=56&sort=1&categoryId=1005",
				page: 1,
			},
			want: "https://www.ximalaya.com/revision/category/v2/albums?pageNum=1&pageSize=56&sort=1&categoryId=1005",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatXimalayaUrl(tt.args.url, tt.args.page); got != tt.want {
				t.Errorf("formatXimalayaUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
