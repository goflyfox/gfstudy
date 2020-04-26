package main

import (
	_ "gfcli/boot"
	_ "gfcli/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
