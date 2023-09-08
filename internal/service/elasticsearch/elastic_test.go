package elasticsearch

import (
	"os"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/genv"
)

func TestInitES(t *testing.T) {
	var (
		ctx = gctx.New()
	)
	os.Setenv("env", "dev")
	genv.Set("GF_GCFG_FILE", "config.dev.yaml")
	InitES(ctx)
}
