package cache

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)


func TestSetCache(t *testing.T) {
    var ctx = gctx.New()
    InitCache(ctx)
    SetCache(ctx, "test_key", "test_value", 60*60)
    cacheValue, err := GetCache(ctx, "test_key")
    if err != nil {
        t.Fatal(err)
    }
    t.Logf("The cacheValue is %+v", cacheValue)

    expireTime, err := defaultCache.GetExpire(ctx, "test_key")
    if err != nil {
        t.Fatal(err)
    }
    t.Logf("The expireTime is %+v", expireTime)
}
