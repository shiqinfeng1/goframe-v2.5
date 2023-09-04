package main

import (
	_ "goframe-v2.5/app/user-manager/internal/packed"

	_ "goframe-v2.5/app/user-manager/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe-v2.5/app/user-manager/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
