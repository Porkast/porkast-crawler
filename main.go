package main

import (
	_ "guoshao-fm-crawler/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"guoshao-fm-crawler/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
