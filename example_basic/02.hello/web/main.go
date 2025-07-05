package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("Welcome GoFrame!!!")
	})
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Writeln("Hello World!")
	})

	s.SetPort(8199)
	s.Run()
}
