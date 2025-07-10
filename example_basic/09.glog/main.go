package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

func main() {
	ctx := gctx.New()
	// 对应默认配置项 logger，默认default
	g.Log().Debug(ctx, "[default]Debug")
	g.Log().Info(ctx, "[default]info")
	g.Log().Warning(ctx, "[default]Warning")
	g.Log().Error(ctx, "[default]Error")
	// 对应 logger.logger1 配置项
	g.Log("logger1").Debug(ctx, "[logger1]Debug")
	g.Log("logger1").Info(ctx, "[logger1]info")
	g.Log("logger1").Warning(ctx, "[logger1]Warning")
	g.Log("logger1").Error(ctx, "[logger1]Error")
	// 对应 logger.logger2 配置项
	g.Log("logger2").Debug(ctx, "[logger2]Debug")
	g.Log("logger2").Info(ctx, "[logger2]info")
	g.Log("logger2").Warning(ctx, "[logger2]Warning")
	g.Log("logger2").Error(ctx, "[logger2]Error")

	// 日志级别设置，过滤掉Info日志信息
	l := glog.New()
	l.Info(ctx, "info1")
	_ = l.SetLevelStr("PROD")
	l.Info(ctx, "info2")
	l.Warning(ctx, "warn3")
	// 支持哪些级别
	// LEVEL_DEBU | LEVEL_INFO | LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT

	// 异常
	g.Log().Panic(ctx, "this is panic！")
	g.Log().Info(ctx, "............")

}
