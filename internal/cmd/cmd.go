package cmd

import (
	"context"
	"guoshao-fm-crawler/internal/service/celery"
	"guoshao-fm-crawler/internal/service/celery/jobs"

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
	celery.RegisterWorker()
	celery.GetClient().StartWorker()
	jobs.StartXiMaLaYaJobs(ctx)
}

func hold() {
	select {}
}
