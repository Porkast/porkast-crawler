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

func StartFirstoryJob(ctx context.Context) {
	_, err := gcron.Add(ctx, consts.PODCAST_WEB_CRAWLER_CRON_PATTERN, func(ctx context.Context) {
		var (
			randomSleepTime time.Duration
		)
		randomSleepTime = getRandomStartTime()
		g.Log().Line().Info(ctx, "start FIRSTORY FM entry jobs, sleep random time : ", randomSleepTime)
		time.Sleep(randomSleepTime)
		if !isJobStarted(ctx, consts.FIRSTORY_ENTRY_JOB) {
			jobIsStarted(ctx, consts.FIRSTORY_ENTRY_JOB)
			AssignFirstoryEntryJob(ctx)
		}
	})

	if err != nil {
		g.Log().Line().Error(ctx, "Add FIRSTORY FM entry jobs cron job failed : ", err)
	}
}

func AssignFirstoryEntryJob(ctx context.Context) {

	var err error

	_, err = celery.GetClient().Delay(consts.FIRSTORY_ENTRY_JOB)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign FIRSTORY_CATEGORY_JOB with failed"))
	}
}

func AssignFirstoryCategoryJob(ctx context.Context, categoryId string, skip int) {
	var err error

	_, err = celery.GetClient().Delay(consts.FIRSTORY_CATEGORY_LIST_JOB, categoryId, skip)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign FIRSTORY_CATEGORY_JOB with failed"))
	}

}

func AssignFirstoryShowRSSJob(ctx context.Context, showId string) {
	var err error

	_, err = celery.GetClient().Delay(consts.FIRSTORY_CATEGORY_SHOW_RSS_JOB, showId)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign FIRSOTRY_SHOW_INFO_JOB with failed"))
	}
}
