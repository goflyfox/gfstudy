package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gsession"
)

const SessionUser = "SessionUser"

func main() {
	s := g.Server()

	// 设置存储方式
	sessionStorage := g.Config().MustGet(gctx.New(), "SessionStorage").String()
	if sessionStorage == "redis" {
		s.SetConfigWithMap(g.Map{
			"SessionIdName":  g.Config().MustGet(gctx.New(), "server.SessionIdName").String(),
			"SessionStorage": gsession.NewStorageRedis(g.Redis()),
		})
	} else if sessionStorage == "memory" {
		s.SetConfigWithMap(g.Map{
			"SessionStorage": gsession.NewStorageMemory(),
		})
	}

	// 常规注册
	group := s.Group("/")
	group.GET("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("index.html", g.Map{
			"title": "登录页面",
		})
	})
	group.POST("/login", func(r *ghttp.Request) {
		username := r.Get("username").String()
		password := r.Get("password").String()

		//dbUsername := "admin"
		//dbPassword := "123456"
		dbUsername := g.Config().MustGet(r.GetCtx(), "username").String()
		dbPassword := g.Config().MustGet(r.GetCtx(), "password").String()
		if username == dbUsername && password == dbPassword {
			// 添加session
			r.Session.Set(SessionUser, g.Map{
				"username": dbUsername,
				"name":     "管理员",
			})
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

	// 用户组
	userGroup := s.Group("/user")
	userGroup.Middleware(MiddlewareAuth)
	// 列表页面
	userGroup.GET("/index", func(r *ghttp.Request) {
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
	userGroup.POST("/logout", func(r *ghttp.Request) {
		// 删除session
		r.Session.Remove(SessionUser)

		r.Response.WriteJson(g.Map{
			"code": 0,
			"msg":  "登出成功",
		})
	})

	s.Run()
}

// 认证中间件
func MiddlewareAuth(r *ghttp.Request) {
	if ok, _ := r.Session.Contains(SessionUser); ok {
		r.Middleware.Next()
	} else {
		// 获取用错误码
		r.Response.WriteJson(g.Map{
			"code": 403,
			"msg":  "您访问超时或已登出",
		})
	}
}
