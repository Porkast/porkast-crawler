package cmd

import (
	"context"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/cache"
	"guoshao-fm-crawler/internal/service/celery"
	"guoshao-fm-crawler/internal/service/celery/jobs"
	"guoshao-fm-crawler/internal/service/celery/worker"
	"guoshao-fm-crawler/internal/service/elasticsearch"
	"os"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/genv"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start podcast crawler",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			initConfig()
			initComponent(ctx)
			initCelery(ctx)
			hold()
			return
		},
	}
)

func initConfig() {
	if os.Getenv("env") == "dev" {
		genv.Set("GF_GCFG_FILE", "config.dev.yaml")
	} else if os.Getenv("env") == "prod" {
		genv.Set("GF_GCFG_FILE", "config.prod.yaml")
	} else {
		genv.Set("GF_GCFG_FILE", "config.yaml")
    }
}

func initComponent(ctx context.Context) {
	cache.InitCache(ctx)
	elasticsearch.InitES(ctx)
}

func initCelery(ctx context.Context) {
	celery.InitCeleryClient(ctx)
	RegisterCeleryWorker()
	celery.GetClient().StartWorker()
	jobs.StartFeedUpdatJobs(ctx)
	jobs.StartXiMaLaYaJobs(ctx)
	jobs.StartLizhiJob(ctx)
	jobs.StartFirstoryJob(ctx)
	jobs.StartSpreakerJob(ctx)
	jobs.StartPodbeanJob(ctx)
	jobs.StartApplePodcastJob(ctx)
}

func RegisterCeleryWorker() {
	// Channel Update Job
	celery.GetClient().Register(consts.CHANNEL_UPDATE_BY_FEED_LINK, worker.ChannelUpdateByFeedLink)
	// XIMALAYA
	celery.GetClient().Register(consts.XIMALAYA_PODCAST_WORKER, worker.ParseXiMaLaYaPodcast)
	celery.GetClient().Register(consts.XIMALAYA_ENTRY_WORKER, worker.ParseXiMaLaYaEntry)
	// LIZHI FM
	celery.GetClient().Register(consts.LIZHI_ENTRY_WORKER, worker.ParseLizhiAllCategories)
	celery.GetClient().Register(consts.LIZHI_CATEGORY_PARSE_WORKER, worker.ParseLizhiPodcastByCategoryPage)
	celery.GetClient().Register(consts.LIZHI_PODCAST_XML_WORKER, worker.ParseLizhiPodcastRSS)
	// FIRSOTRY FM
	celery.GetClient().Register(consts.FIRSTORY_ENTRY_JOB, worker.ParseFistoryAllCategoryList)
	celery.GetClient().Register(consts.FIRSTORY_CATEGORY_LIST_JOB, worker.ParseFirstoryCategoryItemList)
	celery.GetClient().Register(consts.FIRSTORY_CATEGORY_SHOW_RSS_JOB, worker.ParseFirstoryShowRSS)
	// SPREAKER FM
	celery.GetClient().Register(consts.SPREAKER_ENTRY_JOB, worker.ParseSpreakerAllCategoryList)
	celery.GetClient().Register(consts.SPREAKER_SINGLE_CATEGORY_JOB, worker.ParseSpreakerSingleCategory)
	celery.GetClient().Register(consts.SPREAKER_CATEGORY_SHOW_RSS_JOB, worker.ParseSpreakerShowRSS)
	// PODBEAN FM
	celery.GetClient().Register(consts.PODBEAN_ENTRY_JOB, worker.ParsePodbeanAllcategoryList)
	celery.GetClient().Register(consts.PODBEAN_ALL_CATEGORY_POPULAR_JOB, worker.ParsePodbeancategoryPopularShow)
	celery.GetClient().Register(consts.PODBEAN_RSS_JOB, worker.ParsePodbeanShowRSS)
	// APPLE PODCAST
	celery.GetClient().Register(consts.APPLE_PODCAST_ENTRY_WORK, worker.ParseApplePodcastAllCategories)
	celery.GetClient().Register(consts.APPLE_PODCAST_CATEGORY_ITEM_WORK, worker.ParseApplePodcastCategoryItems)
	celery.GetClient().Register(consts.APPLE_PODCAST_ITEM_RSS_WORK, worker.GetApplePodcastItemRSS)
}
func hold() {
	select {}
}
