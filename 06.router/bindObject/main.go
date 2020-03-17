package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct{}

func (c *Controller) Index(r *ghttp.Request) {
	r.Response.Write("index")
}

func (c *Controller) Show(r *ghttp.Request) {
	r.Response.Write("show")
}

func main() {
	s := g.Server()
	c := new(Controller)
	s.BindObject("/object", c)
	s.SetPort(8199)
	s.Run()
}
