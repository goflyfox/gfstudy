package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("Welcome GoFrame!")
	})
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Writeln("Hello World!")
	})
	s.BindHandler("POST:/test", func(r *ghttp.Request) {
		r.Response.Writeln("name:", r.GetString("name"))
		r.Response.Writeln("age:", r.GetInt("age"))
		r.Response.Writeln("func:test")
	})
	s.BindHandler("POST:/test2", func(r *ghttp.Request) {
		r.Response.Writeln("passport:", r.GetString("passport"))
		r.Response.Writeln("password:", r.GetInt("password"))
		r.Response.Writeln("func:test2")
	})
	s.BindHandler("POST:/test3", func(r *ghttp.Request) {
		r.Response.Writeln("Cookie:", r.Header.Get("Cookie"))
		r.Response.Writeln("func:test3")
	})
	g.Config()

	s.SetPort(80)
	s.Run()
}
