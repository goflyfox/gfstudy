package main

import (
	_ "gfcli/internal/packed"
	"github.com/gogf/gf/v2/os/gres"

	"github.com/gogf/gf/v2/os/gctx"

	"gfcli/internal/cmd"
)

func main() {
	gres.Dump()

	cmd.Main.Run(gctx.GetInitCtx())
}
