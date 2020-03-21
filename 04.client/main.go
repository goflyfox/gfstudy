package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	// 默认路径
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("Welcome GoFrame!")
	})
	// GET带参数
	s.BindHandler("GET:/hello", func(r *ghttp.Request) {
		r.Response.Writeln("name:", r.GetString("name"))
		r.Response.Writeln("Hello World!")
	})
	// POST KV
	s.BindHandler("POST:/test", func(r *ghttp.Request) {
		r.Response.Writeln("name:", r.GetString("name"))
		r.Response.Writeln("age:", r.GetInt("age"))
		r.Response.Writeln("func:test")
	})
	// POST JSON
	s.BindHandler("POST:/test2", func(r *ghttp.Request) {
		r.Response.Writeln("passport:", r.GetString("passport"))
		r.Response.Writeln("password:", r.GetString("password"))
		r.Response.Writeln("func:test2")
	})
	// POST Header
	s.BindHandler("POST:/test3", func(r *ghttp.Request) {
		r.Response.Writeln("Cookie:", r.Header.Get("Cookie"))
		r.Response.Writeln("func:test3")
	})
	// POST Header
	s.BindHandler("POST:/test4", func(r *ghttp.Request) {
		r.Response.Writeln("accept-encoding:", r.Header.Get("accept-encoding"))
		r.Response.Writeln("accept-language:", r.Header.Get("accept-encoding"))
		r.Response.Writeln("referer:", r.Header.Get("accept-encoding"))
		r.Response.Writeln("cookie:", r.Header.Get("cookie"))
		r.Response.Writeln("user-agent::test3")
	})

	s.SetPort(80)
	s.Run()
}
