package jobs

import (
	"context"
	"porkast-crawler/internal/service/cache"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/util/grand"
)

func randInt(min int, max int) int {
	return grand.N(min, max)
}

func getRandomStartTime() (startTime time.Duration) {
	var (
		randomInt int
	)

	randomInt = randInt(5, 20)

	startTime = time.Second * time.Duration(randomInt)

	return
}

func isJobStarted(ctx context.Context, key string) (isStart bool) {
	var (
		valueVal *gvar.Var
		err      error
	)

	valueVal, err = cache.GetCache(ctx, key)
	if err != nil {
		isStart = true
	} else if !valueVal.IsEmpty() {
		isStart = true
	} else {
		isStart = false
	}

	return
}

func jobIsStarted(ctx context.Context, key string) {
	cache.SetCache(ctx, key, key, int(60*60))
}
