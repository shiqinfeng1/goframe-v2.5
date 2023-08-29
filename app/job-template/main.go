package main

import (
	_ "goframe-v2.5/app/job-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe-v2.5/app/job-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
