package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	// ##############函数注册/handler start...
	// hello方法，post调用
	s.BindHandler("POST:/handler", func(r *ghttp.Request) {
		g.Log().Info(r.GetCtx(), r.Router.Uri)
		r.Response.Writeln("url:" + r.Router.Uri)
	})
	// 所有方法，url包含name参数
	s.BindHandler("/handler/:name", func(r *ghttp.Request) {
		g.Log().Info(r.GetCtx(), r.Router.Uri)
		// 获取URL name参数
		r.Response.Writeln("name:" + r.Get("name").String())
		r.Response.Writeln("url:" + r.Router.Uri)
	})
	// 所有方法，url包含name参数
	s.BindHandler("/handler/:name/update", func(r *ghttp.Request) {
		g.Log().Info(r.GetCtx(), r.Router.Uri)
		r.Response.Writeln("name:" + r.Get("name").String())
		r.Response.Writeln("url:" + r.Router.Uri)
	})
	// 所有方法，url包含name和action参数
	s.BindHandler("/handler/:name/:action", func(r *ghttp.Request) {
		g.Log().Info(r.GetCtx(), r.Router.Uri)
		r.Response.Writeln("name:" + r.Get("name").String())
		r.Response.Writeln("action:" + r.Get("action").String())
		r.Response.Writeln("url:" + r.Router.Uri)
	})

	// 所有方法，url包含field属性
	s.BindHandler("/handler/list/{field}.html", func(r *ghttp.Request) {
		g.Log().Info(r.GetCtx(), r.Router.Uri)
		// 获取URL field属性
		r.Response.Writeln("field:" + r.Get("field").String())
		r.Response.Writeln("url:" + r.Router.Uri)
	})

	// 方法注册
	s.BindHandler("/handler/total", Total)
	// ##############函数注册 end

	// ##############对象注册/object start...
	// 对象注册
	c := new(Controller)
	s.BindObject("POST:/object", c)
	//  ##############对象注册 end

	// ##############分组注册/group start...
	s.Group("/group", func(group *ghttp.RouterGroup) {
		group.ALL("/all", func(r *ghttp.Request) {
			r.Response.Writeln("all")
		})
		group.GET("/get", func(r *ghttp.Request) {
			r.Response.Writeln("get")
		})
		group.POST("/post", func(r *ghttp.Request) {
			r.Response.Writeln("post")
		})
	})
	//  ##############分组注册 end

	// ##############规范注册/user start...
	s.Group("/user", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			new(Hello),
		)
	})
	//  ##############规范注册 end

	s.Run()
}

var (
	total = gtype.NewInt()
)

func Total(r *ghttp.Request) {
	g.Log().Info(r.GetCtx(), r.Router.Uri)
	r.Response.Write("total:", total.Add(1))
}

// 对象注册
type Controller struct{}

func (c *Controller) Index(r *ghttp.Request) {
	g.Log().Info(r.GetCtx(), r.Router.Uri)
	r.Response.Write("index")
}

func (c *Controller) Show(r *ghttp.Request) {
	g.Log().Info(r.GetCtx(), r.Router.Uri)
	r.Response.Write("show")
}

// 规范注册
type HelloReq struct {
	g.Meta `path:"/hello" method:"get"`
	Name   string `v:"required" dc:"Your name"`
}
type HelloRes struct {
	Reply string `dc:"Reply content"`
}

type Hello struct{}

func (Hello) Say(ctx context.Context, req *HelloReq) (res *HelloRes, err error) {
	g.Log().Info(ctx, ghttp.RequestFromCtx(ctx).Router.Uri)
	g.Log().Debugf(ctx, `receive say: %+v`, req)
	res = &HelloRes{
		Reply: fmt.Sprintf(`Hi %s`, req.Name),
	}
	return
}
