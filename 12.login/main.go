package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	// 常规注册
	group := s.Group("/")
	// 登录页面
	group.GET("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("index.html", g.Map{
			"title": "登录页面",
		})
	})
	// 登录接口
	group.POST("/login", func(r *ghttp.Request) {
		username := r.GetString("username")
		password := r.GetString("password")

		//dbUsername := "admin"
		//dbPassword := "123456"
		dbUsername := g.Config().GetString("username")
		dbPassword := g.Config().GetString("password")
		if username == dbUsername && password == dbPassword {
			r.Response.WriteJson(g.Map{
				"code": 0,
				"msg":  "登录成功",
			})
			r.Exit()
		}

		r.Response.WriteJson(g.Map{
			"code": -1,
			"msg":  "登录失败",
		})
	})
	// 列表页面
	group.GET("/user/index", func(r *ghttp.Request) {
		r.Response.WriteTpl("user_index.html", g.Map{
			"title": "列表页面",
			"dataList": g.List{
				g.Map{
					"date":    "2020-04-01",
					"name":    "朱元璋",
					"address": "江苏110号",
				},
				g.Map{
					"date":    "2020-04-02",
					"name":    "徐达",
					"address": "江苏111号",
				},
				g.Map{
					"date":    "2020-04-03",
					"name":    "李善长",
					"address": "江苏112号",
				},
			}})
	})
	// 登出接口
	group.POST("/logout", func(r *ghttp.Request) {
		r.Response.WriteJson(g.Map{
			"code": 0,
			"msg":  "登出成功",
		})
	})

	s.Run()
}
