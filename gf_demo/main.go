package main

import (
	_ "gf_demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf_demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
