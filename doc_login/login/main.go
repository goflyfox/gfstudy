package main

import (
	_ "gf-login/boot"
	_ "gf-login/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
