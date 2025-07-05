package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	// 默认路径
	s.BindHandler("/", func(r *ghttp.Request) {
		name, err := g.Cfg().Get(r.GetCtx(), "name")
		if err != nil {
			r.Response.Writeln(err.Error())
			r.Exit()
		}
		r.Response.Writeln("配置", name)
		r.Response.Writeln("Welcome GoFrame!")
	})

	s.SetPort(8199)
	s.Run()

}
