package jobs

import (
	"context"
	"fmt"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/celery"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func StartFirstoryJob(ctx context.Context) {
	go func(ctx context.Context) {
		var (
			refreshTime     = time.Hour * 1
			randomSleepTime time.Duration
		)

		for {
			randomSleepTime = getRandomStartTime()
			g.Log().Info(ctx, "start FIRSTORY FM entry jobs, sleep random time : ", randomSleepTime)
			time.Sleep(randomSleepTime)
			if !isJobStarted(ctx, consts.FIRSTORY_ENTRY_JOB) {
				jobIsStarted(ctx, consts.FIRSTORY_ENTRY_JOB)
				AssignLizhiEntryJob(ctx)
			} else {
				g.Log().Info(ctx, "The FIRSTORY FM entry jobs is started, sleep ", refreshTime, " hour")
			}
			time.Sleep(refreshTime)
		}
	}(ctx)
}

func AssignFirstoryEntryJob(ctx context.Context) {

	var err error

	_, err = celery.GetClient().Delay(consts.FIRSTORY_CATEGORY_LIST_JOB)
	if err != nil {
		g.Log().Error(ctx, fmt.Sprintf("Assign FIRSTORY_CATEGORY_JOB with failed"))
	}
}

func AssignFirstoryCategoryJob(ctx context.Context, categoryId string, skip int) {
	var err error

	_, err = celery.GetClient().Delay(consts.FIRSTORY_CATEGORY_LIST_JOB, categoryId, skip)
	if err != nil {
		g.Log().Error(ctx, fmt.Sprintf("Assign FIRSTORY_CATEGORY_JOB with failed"))
	}

}

func AssignFirstoryShowRSSJob(ctx context.Context, showId string) {
	var err error

	_, err = celery.GetClient().Delay(consts.FIRSTORY_CATEGORY_SHOW_RSS_JOB, showId)
	if err != nil {
		g.Log().Error(ctx, fmt.Sprintf("Assign FIRSOTRY_SHOW_INFO_JOB with failed"))
	}
}
