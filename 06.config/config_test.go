package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"testing"
)

// 基本配置使用
func TestConfig(t *testing.T) {
	// 默认当前路径或者config路径，默认文件config.toml
	// /home/www/templates/
	fmt.Println(g.Config().Get("viewpath"))
	fmt.Println(g.Cfg().Get("viewpath"))
	// 127.0.0.1:6379,1
	c := g.Cfg()
	// 分组方式
	fmt.Println(c.Get("redis.cache"))
	// 数组方式：test2
	fmt.Println(c.Get("database.default.1.name"))
}

// 设置路径
func TestConfig2(t *testing.T) {
	// 设置加载文件，默认name为default
	// 设置路径
	g.Cfg().SetPath("configTest")
	// 设置加载文件
	g.Cfg().SetFileName("config1.toml")

	// 打印测试
	fmt.Println(g.Cfg().Get("viewpath"))
	fmt.Println(g.Cfg().Get("study"))
	fmt.Println(g.Cfg().Get("study1"))
	fmt.Println(g.Cfg().Get("config2"))

	// 新的name就是新的实例
	g.Cfg("name").SetPath("configTest")
	g.Cfg("name").SetFileName("config2.toml")
	fmt.Println(g.Cfg("name").Get("viewpath"))
	fmt.Println(g.Cfg("name").Get("study"))
	fmt.Println(g.Cfg("name").Get("study1"))
	fmt.Println(g.Cfg("name").Get("config2"))
}
