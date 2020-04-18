package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gsession"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

const SessionUser = "SessionUser"

func main() {
	s := g.Server()

	// 设置存储方式
	sessionStorage := g.Config().GetString("SessionStorage")
	if sessionStorage == "redis" {
		s.SetSessionStorage(gsession.NewStorageRedis(g.Redis()))
		s.SetSessionIdName(g.Config().GetString("server.SessionIdName"))
	} else if sessionStorage == "memory" {
		s.SetSessionStorage(gsession.NewStorageMemory())
	}

	// 常规注册
	group := s.Group("/")
	group.GET("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("index.html", g.Map{
			"title": "登录页面",
		})
	})

	type User struct {
		Username string `gvalid:"username     @required|length:6,30#请输入用户名称|用户名称长度非法"`
		Password string `gvalid:"password     @required|length:6,30#请输入密码|密码长度非法"`
	}

	group.POST("/login", func(r *ghttp.Request) {
		username := r.GetString("username")
		password := r.GetString("password")

		// 使用结构体定义的校验规则和错误提示进行校验
		if e := gvalid.CheckStruct(User{username, password}, nil); e != nil {
			r.Response.WriteJson(g.Map{
				"code": -1,
				"msg":  e.Error(),
			})
			r.Exit()
		}

		record, err := g.DB().Table("sys_user").Where("login_name = ? ", username).One()
		// 查询数据库异常
		if err != nil {
			r.Response.WriteJson(g.Map{
				"code": -1,
				"msg":  err.Error(),
			})
			r.Exit()
		}

		// 直接存入前端传输的
		if password == record["password"].String() {
			// 添加session
			r.Session.Set(SessionUser, g.Map{
				"username": username,
				"realName": record["real_name"].String(),
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
		realName := gconv.String(r.Session.GetMap(SessionUser)["realName"])
		r.Response.WriteTpl("user_index.html", g.Map{
			"title":    "用户信息列表页面",
			"realName": realName,
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

	// 生成秘钥文件
	// openssl genrsa -out server.key 2048
	// 生成证书文件
	// openssl req -new -x509 -key server.key -out server.crt -days 365
	s.EnableHTTPS("config/server.crt", "config/server.key")
	s.SetHTTPSPort(8080)
	s.SetPort(8199)

	s.Run()
}

// 认证中间件
func MiddlewareAuth(r *ghttp.Request) {
	if r.Session.Contains(SessionUser) {
		r.Middleware.Next()
	} else {
		// 获取用错误码
		r.Response.WriteJson(g.Map{
			"code": 403,
			"msg":  "您访问超时或已登出",
		})
	}
}
