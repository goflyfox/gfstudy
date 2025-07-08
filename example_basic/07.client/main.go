package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	s.Group("/api", func(group *ghttp.RouterGroup) {
		// 默认路径
		group.ALL("/", func(r *ghttp.Request) {
			r.Response.Writeln("Welcome GoFrame!")
		})
		// GET带参数
		group.GET("/get", func(r *ghttp.Request) {
			r.Response.Writeln("func:/get")
			r.Response.Writeln("Hello World!")
			r.Response.Writeln("name:", r.Get("name").String())
		})
		// POST KV
		group.POST("/post", func(r *ghttp.Request) {
			r.Response.Writeln("func:/post")
			r.Response.Writeln("name:", r.Get("name").String())
			r.Response.Writeln("age:", r.Get("age").Int())
		})
		// POST JSON
		group.POST("/post/json", func(r *ghttp.Request) {
			r.Response.Writeln("func:/post/json")
			r.Response.Writeln("passport:", r.Get("passport").String())
			r.Response.Writeln("password:", r.Get("password").String())
		})
		// POST Header
		group.POST("/post/header", func(r *ghttp.Request) {
			r.Response.Writeln("func:/post/header")
			r.Response.Writeln("Cookie:", r.Header.Get("Cookie"))
		})
		// POST Header
		group.POST("/post/header/raw", func(r *ghttp.Request) {
			r.Response.Writeln("func:/post/header/raw")
			h := r.Header
			r.Response.Writeln("accept-encoding:", h.Get("accept-encoding"))
			r.Response.Writeln("accept-language:", h.Get("accept-language"))
			r.Response.Writeln("referer:", h.Get("referer"))
			r.Response.Writeln("cookie:", h.Get("cookie"))
			r.Response.Writeln(r.Cookie.Map())
		})
	})

	s.Run()
}
