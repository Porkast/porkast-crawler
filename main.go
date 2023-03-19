package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "guoshao-fm-crawler/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"guoshao-fm-crawler/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
