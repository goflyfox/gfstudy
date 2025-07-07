package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()
	// 测试日志
	s.BindHandler("/welcome", func(r *ghttp.Request) {
		g.Log().Info(r.GetCtx(), "你来了！")
		r.Response.WriteJson(g.Map{
			"name": r.Get("name").String(),
			"age":  r.Get("age").Int(),
		})
	})
	// 异常处理
	s.BindHandler("/panic", func(r *ghttp.Request) {
		g.Log().Panic(r.GetCtx(), "123")
	})
	// post请求
	s.BindHandler("POST:/hello", func(r *ghttp.Request) {
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		if err := r.Parse(&user); err != nil {
			r.Response.Write(err.Error())
			return
		}
		r.Response.WriteJson(user)
	})
	// 输出：HTML
	s.BindHandler("/index1", func(r *ghttp.Request) {
		r.Response.Writeln("<h1>欢迎收看GoFrame基础教程</h1><p>努力就有收获</p>")
	})
	// 输出：字符串
	s.BindHandler("/index2", func(r *ghttp.Request) {
		r.Response.Writeln("欢迎收看GoFrame基础教程")
	})
	// 输出：使用模版输出HTML
	s.BindHandler("/index3", func(r *ghttp.Request) {
		// 模版输出
		tplContent := `<h1>{{.title}}</h1><p>{{.content}}</p>`
		err := r.Response.WriteTplContent(tplContent, g.Map{
			"title":   "欢迎收看GoFrame基础教程",
			"content": "努力就有收获",
		})
		if err != nil {
			r.Response.Write(err)
		}
	})
	// 输出：XML
	s.BindHandler("/index4", func(r *ghttp.Request) {
		r.Response.WriteXml(g.Map{
			"name": r.Get("name").String(),
			"age":  r.Get("age").Int(),
		})
	})

	s.Run()
}
