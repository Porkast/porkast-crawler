package worker

import (
	"testing"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gfile"
)

func Test_getXimalayaAlbumUrlList(t *testing.T) {
	var (
		sampleJson   *gjson.Json
		albumUrlList []string
		err          error
	)

	sampleJson, err = gjson.Load("./testdata/ximalaya_resp.json")
	if err != nil {
		t.Fatal("Parse ximalaya category response as json failed : ", err)
	}

	albumUrlList = getXimalayaAlbumUrlList(*sampleJson)
	if len(albumUrlList) == 0 {
		t.Fatal("Parse ximalaya album list is empty")
	}

}

func Test_isXimalayaRespXml(t *testing.T) {
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
			name: "Response string is not RSS case",
			args: args{
				respStr: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isXimalayaRespXml(tt.args.respStr); got != tt.want {
				t.Errorf("isXimalayaRespXml() = %v, want %v", got, tt.want)
			}
		})
	}
}
