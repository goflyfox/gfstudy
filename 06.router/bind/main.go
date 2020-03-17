package main

import (
	"github.com/gogf/gf/container/gtype"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var (
	total = gtype.NewInt()
)

func Total(r *ghttp.Request) {
	r.Response.Write("total:", total.Add(1))
}

type Controller struct {
	total *gtype.Int
}

func (c *Controller) Total(r *ghttp.Request) {
	r.Response.Write("total:", c.total.Add(1))
}

func main() {
	s := g.Server()
	// 方法注册
	s.BindHandler("/total", Total)

	// 对象注册
	c := &Controller{
		total: gtype.NewInt(),
	}
	s.BindHandler("/total2", c.Total)
	s.SetPort(8199)
	s.Run()
}
