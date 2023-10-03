package jobs

import (
	"context"
	"fmt"
	"porkast-crawler/internal/consts"
	"porkast-crawler/internal/service/celery"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

func StartSpreakerJob(ctx context.Context) {
	_, err := gcron.Add(ctx, consts.PODCAST_WEB_CRAWLER_CRON_PATTERN, func(ctx context.Context) {
		var (
			randomSleepTime time.Duration
		)
		randomSleepTime = getRandomStartTime()
		g.Log().Line().Info(ctx, "start SPREAKER FM entry jobs, sleep random time : ", randomSleepTime)
		time.Sleep(randomSleepTime)
		if !isJobStarted(ctx, consts.SPREAKER_ENTRY_JOB) {
			jobIsStarted(ctx, consts.SPREAKER_ENTRY_JOB)
			AssignSpreakerEntryJob(ctx)
		}
	})

	if err != nil {
		g.Log().Line().Error(ctx, "Add SPREAKER FM entry jobs cron job failed : ", err)
	}
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
