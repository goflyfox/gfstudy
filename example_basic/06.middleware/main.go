package main

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type HelloReq struct {
	g.Meta `path:"/hello" method:"get"`
}
type HelloRes struct {
}

type Hello struct{}

func (Hello) Say(ctx context.Context, req *HelloReq) (res *HelloRes, err error) {
	g.Log().Info(ctx, "中")
	return
}

func RequestHandle1(r *ghttp.Request) {
	g.Log().Info(r.GetCtx(), "前1")
	r.Middleware.Next()
}

func RequestHandle2(r *ghttp.Request) {
	g.Log().Info(r.GetCtx(), "前2")
	r.Middleware.Next()
}

func RequestHandle3(r *ghttp.Request) {
	g.Log().Info(r.GetCtx(), "前3")
	r.Middleware.Next()
}

func RequestHandle4(r *ghttp.Request) {
	g.Log().Info(r.GetCtx(), "前4")
	r.Middleware.Next()
}

func ResponseHandle1(r *ghttp.Request) {
	r.Middleware.Next()
	g.Log().Info(r.GetCtx(), "后1")
}

func ResponseHandle2(r *ghttp.Request) {
	r.Middleware.Next()
	g.Log().Info(r.GetCtx(), "后2")
}

func ResponseHandle3(r *ghttp.Request) {
	r.Middleware.Next()
	g.Log().Info(r.GetCtx(), "后3")
}

func ResponseHandle4(r *ghttp.Request) {
	r.Middleware.Next()
	g.Log().Info(r.GetCtx(), "后4")
}

func main() {
	s := g.Server()
	s.Use(ghttp.MiddlewareHandlerResponse)
	s.Group("/", func(group *ghttp.RouterGroup) {
		// 前置中间件
		group.Middleware(RequestHandle1)
		group.Middleware(RequestHandle2)

		// 后置中间件
		group.Middleware(ResponseHandle1)
		group.Middleware(ResponseHandle2)

		group.Group("/sub", func(group *ghttp.RouterGroup) {
			// 前置中间件
			group.Middleware(RequestHandle3)
			group.Middleware(RequestHandle4)

			// 后置中间件
			group.Middleware(ResponseHandle3)
			group.Middleware(ResponseHandle4)

			group.Bind(new(Hello))
		})
	})
	s.Run()
}
