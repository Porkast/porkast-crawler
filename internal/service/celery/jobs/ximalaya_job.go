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
		consts.XIMALAYA_CATEGORY_URL_HISTORY,
		consts.XIMALAYA_CATEGORY_URL_XIANGSHEN,
		consts.XIMALAYA_CATEGORY_URL_GEREN_CHENGZHANG,
		consts.XIMALAYA_CATEGORY_URL_RENWEN_GUOXUE,
		consts.XIMALAYA_CATEGORY_URL_LIFE,
		consts.XIMALAYA_CATEGORY_URL_HOT_CHANNEL,
		consts.XIMALAYA_CATEGORY_URL_HISTORY_CHANNEL,
		consts.XIMALAYA_CATEGORY_URL_EMOTION_CHANNEL,
		consts.XIMALAYA_CATEGORY_URL_FINACE_CHANNEL,
		consts.XIMALAYA_CATEGORY_URL_SELF_IMPROVMENT_CHANNEL,
		consts.XIMALAYA_CATEGORY_URL_HEALTH_CHANNEL,
		consts.XIMALAYA_CATEGORY_URL_LIFE_CHANNEL,
		consts.XIMALAYA_CATEGORY_URL_MOVIE_CHANNEL,
		consts.XIMALAYA_CATEGORY_URL_BUSINESS_CHANNEL,
		consts.XIMALAYA_CATEGORY_URL_ENGLISH_CHANNEL,
		consts.XIMALAYA_CATEGORY_URL_CHILDREN_GROWTH_CHANNEL,
		consts.XIMALAYA_CATEGORY_URL_TECH_CHANNEL,
		consts.XIMALAYA_CATEGORY_URL_EDU_EXAM_CHANNEL,
		consts.XIMALAYA_CATEGORY_URL_SPORT_CHANNEL,
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
			g.Log().Line().Info(ctx, "start ximalaya jobs, sleep random time : ", randomSleepTime)
			time.Sleep(randomSleepTime)
			if !isJobStarted(ctx, consts.XIMALAYA_ENTRY_JOB) {
				jobIsStarted(ctx, consts.XIMALAYA_ENTRY_JOB)
				AssignXiMaLaYaEntryJob(ctx)
			} else {
				g.Log().Line().Info(ctx, "The SPREAKER FM entry jobs is started, sleep ", refreshTime, " hour")
			}
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
			g.Log().Line().Debug(ctx, "Assign ximalaya entry work with url : ", targetUrl)
			_, err = celery.GetClient().Delay(consts.XIMALAYA_ENTRY_WORKER, targetUrl)
			if err != nil {
				g.Log().Line().Error(ctx, fmt.Sprintf("Assign XIMALAYA_ENTRY_WORKER with url %s failed : %s", url, err))
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
		g.Log().Line().Error(ctx, fmt.Sprintf("Assign XIMALAYA_PODCAST_WORKER with url %s failed : %s", url, err))
	}

}
