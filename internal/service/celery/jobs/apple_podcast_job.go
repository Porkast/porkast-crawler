package jobs

import (
	"context"
	"fmt"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/celery"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

func StartApplePodcastJob(ctx context.Context) {

	_, err := gcron.Add(ctx, consts.PODCAST_WEB_CRAWLER_CRON_PATTERN, func(ctx context.Context) {
		var (
			randomSleepTime time.Duration
		)
		randomSleepTime = getRandomStartTime()
		g.Log().Line().Info(ctx, "start apple podcast entry jobs, sleep random time : ", randomSleepTime)
		time.Sleep(randomSleepTime)
		if !isJobStarted(ctx, consts.APPLE_PODCAST_ENTRY_WORK) {
			jobIsStarted(ctx, consts.APPLE_PODCAST_ENTRY_WORK)
			AssignApplePodcastEntryJob(ctx)
		}
	})

	if err != nil {
		g.Log().Line().Error(ctx, "Add apple podcast entry jobs cron job failed : ", err)
	}
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
