package jobs

import (
	"context"
	"fmt"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/celery"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func StartLizhiJob(ctx context.Context) {
	go func(ctx context.Context) {

		var (
			refreshTime = time.Hour * 1
		)

		for {
			g.Log().Info(ctx, "start lizhi FM entry jobs")
			AssignLizhiEntryJob(ctx)
			time.Sleep(refreshTime)
		}
	}(ctx)
}

func AssignLizhiEntryJob(ctx context.Context) {
	var (
		err error
	)

	_, err = celery.GetClient().Delay(consts.LIZHI_ENTRY_WORKER, consts.LIZHI_FM_ENTRY_URL)
	if err != nil {
		g.Log().Error(ctx, fmt.Sprintf("Assign LIZHI_ENTRY_WORKER with url %s failed : %s", consts.LIZHI_FM_ENTRY_URL, err))
	}
}

func AssignLizhiCateGoryParseJob(ctx context.Context, url string) {

	var (
		err error
	)

	_, err = celery.GetClient().Delay(consts.LIZHI_CATEGORY_PARSE_WORKER, url)
	if err != nil {
		g.Log().Error(ctx, fmt.Sprintf("Assign LIZHI_CATEGORY_PARSE_WORKER with url %s failed : %s", url, err))
	}
}

func AssignLizhiPodcastXmlJob(ctx context.Context, url string) {

	var (
		err error
	)

	_, err = celery.GetClient().Delay(consts.LIZHI_PODCAST_XML_WORKER, url)
	if err != nil {
		g.Log().Error(ctx, fmt.Sprintf("Assign LIZHI_PODCAST_XML_WORKER with url %s failed : %s", url, err))
	}
}
