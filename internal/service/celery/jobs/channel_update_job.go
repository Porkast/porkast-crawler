package jobs

import (
	"context"
	"fmt"
	"porkast-crawler/internal/consts"
	"porkast-crawler/internal/model/entity"
	"porkast-crawler/internal/service/celery"
	"porkast-crawler/internal/service/internal/dao"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

func StartFeedUpdatJobs(ctx context.Context) {

	_, err := gcron.Add(ctx, consts.FEED_UPDATE_CRON_PATTERN, func(ctx context.Context) {
		var (
			randomSleepTime time.Duration
		)
		randomSleepTime = getRandomStartTime()
		g.Log().Line().Info(ctx, "start channel update jobs, sleep random time : ", randomSleepTime)
		time.Sleep(randomSleepTime)
		if !isJobStarted(ctx, consts.CHANNEL_UPDATE_ENTRY_JOB) {
			jobIsStarted(ctx, consts.CHANNEL_UPDATE_ENTRY_JOB)
			AssignChannelUpdateEntryJob(ctx)
		}
	})

	if err != nil {
		g.Log().Line().Error(ctx, "Add feed updated cron job failed : ", err)
	}
}

func AssignChannelUpdateEntryJob(ctx context.Context) {
	var (
		totalCount int
		offset     = 0
		limit      = 100
		err        error
	)

	totalCount, err = dao.GetAllFeedChannelCount(ctx)
	if err != nil {
		return
	}

	g.Log().Line().Debug(ctx, "The channel total count is ", totalCount)
	for offset < totalCount {
		var (
			channelModelList []entity.FeedChannel
		)
		g.Log().Line().Debug(ctx, "Start from offset ", offset)
		channelModelList = dao.GetFeedChannelList(ctx, offset, limit)
		offset = offset + limit
		for _, channelModel := range channelModelList {
			var err error
			_, err = celery.GetClient().Delay(consts.CHANNEL_UPDATE_BY_FEED_LINK, channelModel.FeedLink)
			if err != nil {
				g.Log().Line().Error(ctx, fmt.Sprintf("Assign CHANNEL_UPDATE_BY_FEED_LINK failed : %s\n", err))
			}
		}
	}

}
