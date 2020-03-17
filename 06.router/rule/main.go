package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("POST:/hello", func(r *ghttp.Request) {
		r.Response.Writeln("url" + r.Router.Uri)
	})
	s.BindHandler("/:name", func(r *ghttp.Request) {
		// 获取URL name参数
		r.Response.Writeln("name:" + r.GetString("name"))
		r.Response.Writeln("url" + r.Router.Uri)
	})
	s.BindHandler("/:name/update", func(r *ghttp.Request) {
		r.Response.Writeln("name:" + r.GetString("name"))
		r.Response.Writeln("url" + r.Router.Uri)
	})
	s.BindHandler("/:name/:action", func(r *ghttp.Request) {
		r.Response.Writeln("name:" + r.GetString("name"))
		r.Response.Writeln("action:" + r.GetString("action"))
		r.Response.Writeln("url" + r.Router.Uri)
	})
	s.BindHandler("/user/list/{field}.html", func(r *ghttp.Request) {
		// 获取URL field属性
		r.Response.Writeln("field:" + r.GetString("field"))
		r.Response.Writeln("url" + r.Router.Uri)
	})
	s.SetPort(8199)
	s.Run()
}
