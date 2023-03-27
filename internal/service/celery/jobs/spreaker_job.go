package jobs

import (
	"context"
	"fmt"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/celery"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func StartSpreakerJob(ctx context.Context) {
	go func(ctx context.Context) {

		var (
			refreshTime     = time.Hour * 1
			randomSleepTime time.Duration
		)

		for {
			randomSleepTime = getRandomStartTime()
			g.Log().Line().Info(ctx, "start SPREAKER FM entry jobs, sleep random time : ", randomSleepTime)
			time.Sleep(randomSleepTime)
			if !isJobStarted(ctx, consts.SPREAKER_ENTRY_JOB) {
				jobIsStarted(ctx, consts.SPREAKER_ENTRY_JOB)
				AssignSpreakerEntryJob(ctx)
			} else {
				g.Log().Line().Info(ctx, "The SPREAKER FM entry jobs is started, sleep ", refreshTime, " hour")
			}
			time.Sleep(refreshTime)
		}
	}(ctx)
}

func AssignSpreakerEntryJob(ctx context.Context) {
	var (
		err error
	)
	_, err = celery.GetClient().Delay(consts.SPREAKER_ENTRY_JOB)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign SPREAKER_ENTRY_JOB with failed"))
	}
}

func AssignSpreakerSingleCategoryJob(ctx context.Context, url string) {

	var (
		err error
	)
	_, err = celery.GetClient().Delay(consts.SPREAKER_SINGLE_CATEGORY_JOB, url)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign SPREAKER_SINGLE_CATEGORY_JOB with failed"))
	}
}

func AssignSpreakerShowRSSJob(ctx context.Context, url string) {

	var (
		err error
	)
	_, err = celery.GetClient().Delay(consts.SPREAKER_CATEGORY_SHOW_RSS_JOB, url)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign SPREAKER_SINGLE_CATEGORY_JOB with failed"))
	}
}
