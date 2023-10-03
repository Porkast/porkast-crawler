package worker

import (
	"porkast-crawler/internal/service/network"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func ChannelUpdateByFeedLink(feedLink string) {
	var (
		ctx     = gctx.New()
		respStr string
	)

	respStr = network.TryGetRSSContent(ctx, feedLink)
	if isStringRSSXml(respStr) {
		storeFeed(ctx, respStr, feedLink, "ChannelUpdateByFeedLink")
	} else {
		g.Log().Line().Errorf(ctx, "The response by feed link %s is not RSS XML", feedLink)
	}
}
