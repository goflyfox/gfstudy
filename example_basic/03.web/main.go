package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	// 测试日志
	s.BindHandler("/welcome", func(r *ghttp.Request) {
		g.Log().Info(r.GetCtx(), "你来了！")
		g.Log().Error(r.GetCtx(), "你异常啦！")
		r.Response.Write("哈喽世界！")
	})
	// 异常处理
	s.BindHandler("/panic", func(r *ghttp.Request) {
		g.Log().Panic(r.GetCtx(), "123")
	})
	// post请求
	s.BindHandler("POST:/hello", func(r *ghttp.Request) {
		r.Response.Writeln("Hello World!")
	})
	s.Run()
}
