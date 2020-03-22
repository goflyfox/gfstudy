package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	// 默认路径
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("配置", g.Config().GetString("name"))
		r.Response.Writeln("Welcome GoFrame!")
	})

	s.SetPort(80)
	s.Run()

}
