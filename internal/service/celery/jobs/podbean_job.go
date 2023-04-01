package jobs

import (
	"context"
	"fmt"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/celery"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func StartPodbeanJob(ctx context.Context) {
	go func(ctx context.Context) {

		var (
			refreshTime     = time.Hour * 1
			randomSleepTime time.Duration
		)

		for {
			randomSleepTime = getRandomStartTime()
			g.Log().Line().Info(ctx, "start PODBEAN FM entry jobs, sleep random time : ", randomSleepTime)
			time.Sleep(randomSleepTime)
			if !isJobStarted(ctx, consts.PODBEAN_ENTRY_JOB) {
				jobIsStarted(ctx, consts.PODBEAN_ENTRY_JOB)
				AssignPodbeanEntryJob(ctx)
			} else {
				g.Log().Line().Info(ctx, "The PODBEAN FM entry jobs is started, sleep ", refreshTime, " hour")
			}
			time.Sleep(refreshTime)
		}
	}(ctx)
}

func AssignPodbeanEntryJob(ctx context.Context) {
	var (
		err error
	)
	_, err = celery.GetClient().Delay(consts.PODBEAN_ENTRY_JOB, consts.PODBEAN_ALL_CATEGORY_URL)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign PODBEAN_ENTRY_JOB with failed"))
	}
}

func AssignPodbeanCategoryPopularListJob(ctx context.Context, category string, page int) {
	var (
		err error
	)
	_, err = celery.GetClient().Delay(consts.PODBEAN_ALL_CATEGORY_POPULAR_JOB, category, 1)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign PODBEAN_ALL_CATEGORY_POPULAR_JOB with failed"))
	}
}

func AssignPodbeanRSSJob(ctx context.Context, rssLink string) {
	var (
		err error
	)
	_, err = celery.GetClient().Delay(consts.PODBEAN_RSS_JOB, rssLink)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign PODBEAN_RSS_JOB with failed"))
	}
}
