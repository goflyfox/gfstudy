package main

import (
	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	// 常规注册
	// hello方法，post调用
	s.BindHandler("POST:/hello", func(r *ghttp.Request) {
		r.Response.Writeln("url" + r.Router.Uri)
	})
	// 所有方法，url包含name参数
	s.BindHandler("/:name", func(r *ghttp.Request) {
		// 获取URL name参数
		r.Response.Writeln("name:" + r.GetString("name"))
		r.Response.Writeln("url" + r.Router.Uri)
	})
	// 所有方法，url包含name参数
	s.BindHandler("/:name/update", func(r *ghttp.Request) {
		r.Response.Writeln("name:" + r.GetString("name"))
		r.Response.Writeln("url" + r.Router.Uri)
	})
	// 所有方法，url包含name和action参数
	s.BindHandler("/:name/:action", func(r *ghttp.Request) {
		r.Response.Writeln("name:" + r.GetString("name"))
		r.Response.Writeln("action:" + r.GetString("action"))
		r.Response.Writeln("url" + r.Router.Uri)
	})
	// 所有方法，url包含field属性
	s.BindHandler("/user/list/{field}.html", func(r *ghttp.Request) {
		// 获取URL field属性
		r.Response.Writeln("field:" + r.GetString("field"))
		r.Response.Writeln("url" + r.Router.Uri)
	})

	// 方法注册
	s.BindHandler("/total", Total)

	// 对象注册
	c := new(Controller)
	s.BindObject("POST:/object", c)

	// 分组注册及中间件
	group := s.Group("/api")
	group.Middleware(MiddlewareTest)
	group.ALL("/all", func(r *ghttp.Request) {
		r.Response.Writeln("all")
	})
	group.GET("/get", func(r *ghttp.Request) {
		r.Response.Writeln("get")
	})
	group.POST("/post", func(r *ghttp.Request) {
		r.Response.Writeln("post")
	})

	// request and response
	s.BindHandler("POST:/test", func(r *ghttp.Request) {
		r.Response.WriteJson(g.Map{
			"name": r.GetString("name"),
			"age":  r.GetInt("age"),
			"sex":  r.Header.Get("sex"),
		})
	})

	s.SetPort(8199)
	s.Run()
}

var (
	total = gtype.NewInt()
)

func Total(r *ghttp.Request) {
	r.Response.Write("total:", total.Add(1))
}

// 对象注册
type Controller struct{}

func (c *Controller) Index(r *ghttp.Request) {
	r.Response.Write("index")
}

func (c *Controller) Show(r *ghttp.Request) {
	r.Response.Write("show")
}

// 中间件
func MiddlewareTest(r *ghttp.Request) {
	// 前置逻辑
	r.Response.Writeln("###start")
	r.Middleware.Next()
	// 后置逻辑
	r.Response.Writeln("###end")
}
