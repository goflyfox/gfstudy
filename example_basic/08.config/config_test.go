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
	name, err := g.Cfg().Get(ctx, "name")
	fmt.Println(name, err)
	// /home/www/template/
	fmt.Println(g.Cfg().MustGet(ctx, "viewpath"))
	// 127.0.0.1:6379,1
	// 分组方式
	fmt.Println(g.Cfg().MustGet(ctx, "redis.cache"))
	// 数组方式：test2
	fmt.Println(g.Cfg().MustGet(ctx, "database.default.1.name"))
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
	fmt.Println("viewpath:", g.Cfg().MustGet(ctx, "viewpath"))
	fmt.Println("study:", g.Cfg().MustGet(ctx, "study"))
	fmt.Println("study1:", g.Cfg().MustGet(ctx, "study1"))
	fmt.Println("config2:", g.Cfg().MustGet(ctx, "config2"))

	// 新的name就是新的实例
	nameCfg := g.Cfg("name")
	nameCfg.GetAdapter().(*gcfg.AdapterFile).SetPath("configTest")
	nameCfg.GetAdapter().(*gcfg.AdapterFile).SetFileName("config2.yaml")
	fmt.Println("viewpath:", nameCfg.MustGet(ctx, "viewpath"))
	fmt.Println("study:", nameCfg.MustGet(ctx, "study"))
	fmt.Println("study1:", nameCfg.MustGet(ctx, "study1"))
	fmt.Println("config2:", nameCfg.MustGet(ctx, "config2"))
}
