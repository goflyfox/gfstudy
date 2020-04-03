package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("Welcome GoFrame!!!")
	})
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Writeln("Hello World!")
	})

	s.SetPort(8080)
	s.Run()
}
