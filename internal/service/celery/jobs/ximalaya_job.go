package jobs

import (
	"context"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/celery"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func StartXiMaLaYaJobs(ctx context.Context) {
	go func(ctx context.Context) {
		var (
			refreshTime = time.Hour * 1
		)

		for {
			g.Log().Info(ctx, "start ximalaya jobs")
			AssignXiMaLaYaEntryJob(ctx)
			time.Sleep(refreshTime)
		}
	}(ctx)
}

func AssignXiMaLaYaEntryJob(ctx context.Context) {
	var (
		entryUrl = "ximalaya entry url"
		err      error
	)

	_, err = celery.GetClient().Delay(consts.XIMALAYA_ENTRY_WORKER, entryUrl)
	if err != nil {
		g.Log().Error(ctx, "run XIMALAYA_ENTRY_WORKER failed : ", err)
	}
}

func AssignXimalayaPodcastJob(ctx context.Context, url string) {
	var (
		err error
	)

	_, err = celery.GetClient().Delay(consts.XIMALAYA_PODCAST_WORKER, url)
	if err != nil {
		g.Log().Error(ctx, "run XIMALAYA_PODCAST_WORKER failed : ", err)
	}

}
