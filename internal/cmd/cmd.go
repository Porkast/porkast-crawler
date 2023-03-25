package cmd

import (
	"context"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/cache"
	"guoshao-fm-crawler/internal/service/celery"
	"guoshao-fm-crawler/internal/service/celery/jobs"
	"guoshao-fm-crawler/internal/service/celery/worker"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start podcast crawler",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			initCache(ctx)
			initCelery(ctx)
			hold()
			return
		},
	}
)

func initConfig()  {
	if os.Getenv("env") == "dev" {
		g.Log().SetLevel(glog.LEVEL_ALL)
	}
}

func initCache(ctx context.Context) {
	cache.InitCache(ctx)
}

func initCelery(ctx context.Context) {
	celery.InitCeleryClient(ctx)
	RegisterCeleryWorker()
	celery.GetClient().StartWorker()
	jobs.StartXiMaLaYaJobs(ctx)
	jobs.StartLizhiJob(ctx)
	jobs.StartFirstoryJob(ctx)
}

func RegisterCeleryWorker() {
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
}
func hold() {
	select {}
}
