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

func StartPodbeanJob(ctx context.Context) {
	_, err := gcron.Add(ctx, consts.PODCAST_WEB_CRAWLER_CRON_PATTERN, func(ctx context.Context) {
		var (
			randomSleepTime time.Duration
		)
		randomSleepTime = getRandomStartTime()
		g.Log().Line().Info(ctx, "start PODBEAN FM entry jobs, sleep random time : ", randomSleepTime)
		time.Sleep(randomSleepTime)
		if !isJobStarted(ctx, consts.PODBEAN_ENTRY_JOB) {
			jobIsStarted(ctx, consts.PODBEAN_ENTRY_JOB)
			AssignPodbeanEntryJob(ctx)
		}
	})

	if err != nil {
		g.Log().Line().Error(ctx, "Add PODBEAN FM entry jobs cron job failed : ", err)
	}
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
