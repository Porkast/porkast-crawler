package worker

import (
	"guoshao-fm-crawler/internal/service/network"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func ParseXiMaLaYaPodcast(url string) {
	var (
		ctx        = gctx.New()
		podcastUrl string
		respStr    string
	)
	g.Log().Info(ctx, "start run task with url : ", url)
	podcastUrl = url + ".xml"
	respStr = network.GetContent(ctx, podcastUrl)
    if respStr != "" {
        //The ximalaya album is RSS

    }
}

func ParseXiMaLaYaEntry(url string) {
	var (
		ctx = gctx.New()
	)
	g.Log().Info(ctx, "start run task with url : ", url)
}
