package main

import (
	_ "porkast-crawler/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"porkast-crawler/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
