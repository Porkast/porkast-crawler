package jobs

import (
	"context"
	"fmt"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/celery"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func StartApplePodcastJob(ctx context.Context) {
	go func(ctx context.Context) {
		var (
			refreshTime     = time.Hour * 6
			randomSleepTime time.Duration
		)

		for {
			randomSleepTime = getRandomStartTime()
			g.Log().Line().Info(ctx, "start apple podcast entry jobs, sleep random time : ", randomSleepTime)
			time.Sleep(randomSleepTime)
			if !isJobStarted(ctx, consts.APPLE_PODCAST_ENTRY_WORK) {
				jobIsStarted(ctx, consts.APPLE_PODCAST_ENTRY_WORK)
				AssignApplePodcastEntryJob(ctx)
			} else {
				g.Log().Line().Info(ctx, "The apple podcast entry jobs is started, sleep ", refreshTime, " hour")
			}
			time.Sleep(refreshTime)
		}
	}(ctx)
}

func AssignApplePodcastEntryJob(ctx context.Context) {
	var (
		err error
	)

	_, err = celery.GetClient().Delay(consts.APPLE_PODCAST_ENTRY_WORK, consts.APPLE_PODCAST_ENTRY_URL)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign APPLE_PODCAST_ENTRY_WORK failed : %s", err))
	}
}

func AssignApplePodcastCategoryItemJob(ctx context.Context, categoryUrl string) {
	var (
		err error
	)

	_, err = celery.GetClient().Delay(consts.APPLE_PODCAST_CATEGORY_ITEM_WORK, categoryUrl)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign APPLE_PODCAST_CATEGORY_ITEM_WORK with url %s failed : %s", categoryUrl, err))
	}
}

func AssignApplePodcastItemRSSJob(ctx context.Context, itemUrl string) {
	var (
		err error
	)

	_, err = celery.GetClient().Delay(consts.APPLE_PODCAST_ITEM_RSS_WORK, itemUrl)
	if err != nil {
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign APPLE_PODCAST_ITEM_RSS_WORK with url %s failed : %s", itemUrl, err))
	}
}
