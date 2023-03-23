package jobs

import (
	"context"
	"fmt"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/celery"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func getXimalayaEntryUrlList() (urlList []string) {

	urlList = []string{
		consts.XIMALAYA_CATEGORY_URL,
		consts.XIMALAYA_CATEGORY_URL_HOT,
		consts.XIMALAYA_CATEGORY_URL_MUSIC,
		consts.XIMALAYA_CATEGORY_URL_ENTERTAINENT,
		consts.XIMALAYA_CATEGORY_URL_MUSIC,
		consts.XIMALAYA_CATEGORY_URL_FOREIGN_LANGUAGE,
		consts.XIMALAYA_CATEGORY_URL_CHILDREN,
		consts.XIMALAYA_CATEGORY_URL_BUSINESS,
	}

	return
}

func StartXiMaLaYaJobs(ctx context.Context) {
	go func(ctx context.Context) {
		var (
			refreshTime     = time.Hour * 1
			randomSleepTime time.Duration
		)

		for {
            randomSleepTime = getRandomStartTime()
            g.Log().Info(ctx, "start ximalaya jobs, sleep random time : ", randomSleepTime)
			time.Sleep(randomSleepTime)
			AssignXiMaLaYaEntryJob(ctx)
			time.Sleep(refreshTime)
		}
	}(ctx)
}

func AssignXiMaLaYaEntryJob(ctx context.Context) {
	var (
		totalPage = 50
		err       error
	)

	for _, url := range getXimalayaEntryUrlList() {
		for i := 0; i < totalPage-1; i++ {
			var (
				targetUrl   string
				currentPage int
			)
			currentPage = i + 1
			targetUrl = formatXimalayaUrl(url, currentPage)
			if !isJobStarted(ctx, targetUrl) {
				jobIsStarted(ctx, targetUrl)
				g.Log().Debug(ctx, "Assign ximalaya entry work with url : ", targetUrl)
				_, err = celery.GetClient().Delay(consts.XIMALAYA_ENTRY_WORKER, targetUrl)
				if err != nil {
					g.Log().Error(ctx, fmt.Sprintf("Assign XIMALAYA_ENTRY_WORKER with url %s failed : %s", url, err))
				}
			} else {
                g.Log().Info(ctx, "The ximalaya FM entry jobs is started")
            }
		}
	}

}

func formatXimalayaUrl(url string, page int) string {
	var (
		formatUrl string
	)

	formatUrl = fmt.Sprintf(url, page)

	return formatUrl
}

func AssignXimalayaPodcastJob(ctx context.Context, url string) {
	var (
		err error
	)

	_, err = celery.GetClient().Delay(consts.XIMALAYA_PODCAST_WORKER, url)
	if err != nil {
		g.Log().Error(ctx, fmt.Sprintf("Assign XIMALAYA_PODCAST_WORKER with url %s failed : %s", url, err))
	}

}
