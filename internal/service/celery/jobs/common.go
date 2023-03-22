package jobs

import (
	"context"
	"guoshao-fm-crawler/internal/service/cache"
	"math/rand"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
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
		isStart = false
	} else if valueVal != nil {
		isStart = true
	} else {
		isStart = false
	}

	return
}

func jobIsStarted(ctx context.Context, key string)  {
	cache.SetCache(ctx, key, key, int(time.Minute * 60))
}