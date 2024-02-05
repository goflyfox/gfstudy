package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	group := s.Group("/api")
	// 默认路径
	group.ALL("/", func(r *ghttp.Request) {
		r.Response.Writeln("Welcome GoFrame!")
	})
	// GET带参数
	group.GET("/hello", func(r *ghttp.Request) {
		r.Response.Writeln("Hello World!")
		r.Response.Writeln("name:", r.GetString("name"))
	})
	// POST KV
	group.POST("/test", func(r *ghttp.Request) {
		r.Response.Writeln("func:test")
		r.Response.Writeln("name:", r.GetString("name"))
		r.Response.Writeln("age:", r.GetInt("age"))
	})
	// POST JSON
	group.POST("/test2", func(r *ghttp.Request) {
		r.Response.Writeln("func:test2")
		r.Response.Writeln("passport:", r.GetString("passport"))
		r.Response.Writeln("password:", r.GetString("password"))
	})
	// POST Header
	group.POST("/test3", func(r *ghttp.Request) {
		r.Response.Writeln("func:test3")
		r.Response.Writeln("Cookie:", r.Header.Get("Cookie"))
	})
	// POST Header
	group.POST("/test4", func(r *ghttp.Request) {
		r.Response.Writeln("func:test4")
		h := r.Header
		r.Response.Writeln("accept-encoding:", h.Get("accept-encoding"))
		r.Response.Writeln("accept-language:", h.Get("accept-language"))
		r.Response.Writeln("referer:", h.Get("referer"))
		r.Response.Writeln("cookie:", h.Get("cookie"))
		r.Response.Writeln(r.Cookie.Map())
	})

	s.SetPort(80)
	s.Run()
}
