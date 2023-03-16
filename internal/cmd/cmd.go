package cmd

import (
	"context"
	"guoshao-fm-crawler/internal/consts"
	"guoshao-fm-crawler/internal/service/celery"
	"guoshao-fm-crawler/internal/service/celery/jobs"
	"guoshao-fm-crawler/internal/service/celery/worker"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start podcast crawler",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			initCelery(ctx)
			hold()
			return
		},
	}
)

func initCelery(ctx context.Context) {
	celery.InitCeleryClient(ctx)
	celery.GetClient().StartWorker()
	RegisterCeleryWorker()
	jobs.StartXiMaLaYaJobs(ctx)
}

func RegisterCeleryWorker() {
	celery.GetClient().Register(consts.XIMALAYA_PODCAST_WORKER, worker.ParseXiMaLaYaPodcast)
	celery.GetClient().Register(consts.XIMALAYA_ENTRY_WORKER, worker.ParseXiMaLaYaEntry)
}
func hold() {
	select {}
}
