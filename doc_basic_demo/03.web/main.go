package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	s := g.Server()
	g.Cfg()
	// 测试日志
	s.BindHandler("/welcome", func(r *ghttp.Request) {

		glog.Info(r.GetCtx(), "你来了！")
		glog.Error(r.GetCtx(), "你异常啦！")
		r.Response.Write("哈喽世界！")
	})
	// 异常处理
	s.BindHandler("/panic", func(r *ghttp.Request) {
		glog.Panic(r.GetCtx(), "123")
	})
	// post请求
	s.BindHandler("POST:/hello", func(r *ghttp.Request) {
		r.Response.Writeln("Hello World!")
	})
	s.Run()
}
