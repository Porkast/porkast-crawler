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
			refreshTime     = time.Hour * 6
			randomSleepTime time.Duration
		)

		for {
			randomSleepTime = getRandomStartTime()
			g.Log().Line().Info(ctx, "start lizhi FM entry jobs, sleep random time : ", randomSleepTime)
			time.Sleep(randomSleepTime)
			if !isJobStarted(ctx, consts.LIZHI_ENTRY_WORKER) {
				jobIsStarted(ctx, consts.LIZHI_ENTRY_WORKER)
				AssignLizhiEntryJob(ctx)
			} else {
				g.Log().Line().Info(ctx, "The lizhi FM entry jobs is started, sleep ", refreshTime, " hour")
			}
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
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign LIZHI_ENTRY_WORKER with url %s failed : %s", consts.LIZHI_FM_ENTRY_URL, err))
	}
}

func AssignLizhiCategoryParseJob(ctx context.Context, url string) {

	var (
		err error
	)

	_, err = celery.GetClient().Delay(consts.LIZHI_CATEGORY_PARSE_WORKER, url)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign LIZHI_CATEGORY_PARSE_WORKER with url %s failed : %s", url, err))
	}
}

func AssignLizhiPodcastXmlJob(ctx context.Context, url string) {

	var (
		err error
	)

	_, err = celery.GetClient().Delay(consts.LIZHI_PODCAST_XML_WORKER, url)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign LIZHI_PODCAST_XML_WORKER with url %s failed : %s", url, err))
	}
}
