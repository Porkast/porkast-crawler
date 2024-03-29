package cache

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

var defaultCache *gcache.Cache
var redisClient *gredis.Redis

func InitCache(ctx context.Context) {
	var (
		err error
	)
	defaultCache = gcache.New()
	redisClient = initRedisClient(ctx)
	if redisClient != nil {
		_, err = redisClient.Do(ctx, "SET", "Test", "test_value")
		if err == nil {
			defaultCache.SetAdapter(gcache.NewAdapterRedis(redisClient))
		} else {
			g.Log().Line().Fatal(ctx, "Initial redis failed : ", err)
		}
	}
}

func SetCache(ctx context.Context, key, value string, expireSecond int) error {
	return defaultCache.Set(ctx, key, value, time.Duration(expireSecond)*time.Second)
}

func GetCache(ctx context.Context, key string) (*gvar.Var, error) {
	return defaultCache.Get(ctx, key)
}

func DoHSet(ctx context.Context, key, mapKey string, mapValue g.Map) {
	redisClient.MustDo(ctx, "HSET", key, mapKey, mapValue)
}

func GetHSet(ctx context.Context, key, mapKey string) (value *gvar.Var) {
	value = redisClient.MustDo(ctx, "HGET", key, mapKey)
	return
}
