package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	// 常规注册
	group := s.Group("/")

	// 模板文件
	group.GET("/", func(r *ghttp.Request) {
		r.Response.WriteTpl("index.html", g.Map{
			"title": "列表页面",
			"show": true,
			"listData": g.List{
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

	// 字符串传入
	group.GET("/template", func(r *ghttp.Request) {
		tplContent := `id:${.id}, name:${.name}`
		r.Response.WriteTplContent(tplContent, g.Map{
			"id"   : 123,
			"name" : "john",
		})
	})

	s.Run()
}
