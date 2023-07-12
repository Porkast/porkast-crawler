package jobs

import (
	"context"
	"fmt"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/model/entity"
	"guoshao-fm-crawler/internal/service/celery"
	"guoshao-fm-crawler/internal/service/internal/dao"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func StartFeedUpdatJobs(ctx context.Context) {
	go func(ctx context.Context) {
		var (
			refreshTime     = time.Hour * 3
			randomSleepTime time.Duration
		)

		for {
			randomSleepTime = getRandomStartTime()
			g.Log().Line().Info(ctx, "start channel update jobs, sleep random time : ", randomSleepTime)
			time.Sleep(randomSleepTime)
			if !isJobStarted(ctx, consts.CHANNEL_UPDATE_ENTRY_JOB) {
				jobIsStarted(ctx, consts.CHANNEL_UPDATE_ENTRY_JOB)
				AssignChannelUpdateEntryJob(ctx)
			} else {
				g.Log().Line().Info(ctx, "The channel update entry jobs is started, sleep ", refreshTime, " hour")
			}
			time.Sleep(refreshTime)
		}
	}(ctx)
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
