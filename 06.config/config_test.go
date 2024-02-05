package main

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

// 基本配置使用
func TestConfig(t *testing.T) {
	ctx := gctx.New()
	// 默认当前路径或者config路径，默认文件config.yaml
	// /home/www/template/
	fmt.Println(g.Config().MustGet(ctx, "viewpath"))
	fmt.Println(g.Cfg().MustGet(ctx, "viewpath"))
	// 127.0.0.1:6379,1
	c := g.Cfg()
	// 分组方式
	fmt.Println(c.MustGet(ctx, "redis.cache"))
	// 数组方式：test2
	fmt.Println(c.MustGet(ctx, "database.default.1.name"))
}

// 设置路径
func TestConfig2(t *testing.T) {
	ctx := gctx.New()
	// 设置加载文件，默认name为default
	// 设置路径
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetPath("configTest")
	// 设置加载文件
	g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("config1.yaml")

	// 打印测试
	fmt.Println(g.Cfg().MustGet(ctx, "viewpath"))
	fmt.Println(g.Cfg().MustGet(ctx, "study"))
	fmt.Println(g.Cfg().MustGet(ctx, "study1"))
	fmt.Println(g.Cfg().MustGet(ctx, "config2"))

	// 新的name就是新的实例
	g.Cfg("name").GetAdapter().(*gcfg.AdapterFile).SetPath("configTest")
	g.Cfg("name").GetAdapter().(*gcfg.AdapterFile).SetFileName("config2.yaml")
	fmt.Println(g.Cfg("name").MustGet(ctx, "viewpath"))
	fmt.Println(g.Cfg("name").MustGet(ctx, "study"))
	fmt.Println(g.Cfg("name").MustGet(ctx, "study1"))
	fmt.Println(g.Cfg("name").MustGet(ctx, "config2"))
}
