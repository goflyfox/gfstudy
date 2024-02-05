package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

func main() {
	// 对应默认配置项 logger，默认default
	g.Log().Debug("[default]Debug")
	g.Log().Info("[default]info")
	g.Log().Warning("[default]Warning")
	g.Log().Error("[default]Error")
	// 对应 logger.logger1 配置项
	g.Log("logger1").Debug("[logger1]Debug")
	g.Log("logger1").Info("[logger1]info")
	g.Log("logger1").Warning("[logger1]Warning")
	g.Log("logger1").Error("[logger1]Error")
	// 对应 logger.logger2 配置项
	g.Log("logger2").Debug("[logger2]Debug")
	g.Log("logger2").Info("[logger2]info")
	g.Log("logger2").Warning("[logger2]Warning")
	g.Log("logger2").Error("[logger2]Error")

	// 日志级别设置，过滤掉Info日志信息
	l := glog.New()
	l.Info("info1")
	l.SetLevel(glog.LEVEL_ALL ^ glog.LEVEL_INFO)
	l.Info("info2")
	// 支持哪些级别
	// LEVEL_DEBU | LEVEL_INFO | LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT

	// 异常
	g.Log().Panic("this is panic！")
	g.Log().Info("............")

}
