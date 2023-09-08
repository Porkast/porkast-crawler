package utility

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
)

func TestParseFeed(t *testing.T) {
	type args struct {
		ctx    context.Context
		rssStr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Parse rss string",
			args: args{
				ctx:    gctx.New(),
				rssStr: gfile.GetContents("./testdata/ximalaya_rss_resp.xml"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFeed := ParseFeed(tt.args.ctx, tt.args.rssStr)
			if gotFeed == nil {
				t.Fatal("Parse rss string failed")
			}
		})
	}
}
