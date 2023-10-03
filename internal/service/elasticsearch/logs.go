package elasticsearch

import (
	"context"
	"porkast-crawler/internal/model/entity"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/olivere/elastic/v7"
)

func (c *GSElastic) StoreLogs(ctx context.Context, time, level, content string) {

	esLog := entity.LogESData{
		Id:      gctx.CtxId(ctx),
		Time:    time,
		Level:   level,
		Content: content,
	}

	_, err := elastic.NewIndexService(c.Client).Index("gs_log").Id(esLog.Id).BodyJson(esLog).Do(ctx)
	if err != nil {
		// TODO: log error without update to search engine
	}
}