package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/guuid"
)

func Login(r *ghttp.Request) {
	if r.GetString("username") == "admin" &&
		r.GetString("password") == "123456" {
		r.Response.WriteJson(g.Map{
			"code": 0,
			"msg":  "登录成功",
			"data": guuid.New().String(),
		})
		r.Exit()
	}

	r.Response.WriteJson(g.Map{
		"code": -1,
		"msg":  "登录失败",
		"data": "",
	})

}
