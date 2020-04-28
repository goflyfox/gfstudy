package main

import (
	_ "gfcli/boot"
	_ "gfcli/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gres"
)

func main() {
	gres.Dump()
	g.Server().Run()
}
