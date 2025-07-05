package main

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.New()
	// redis字符串操作
	g.Redis().Set(ctx, "k", "v")
	v, _ := g.Redis().Get(ctx, "k")
	g.Log().Info(ctx, v.String())

	// 获取cache链接,DoVar转换
	v2, _ := g.Redis("cache").Get(ctx, "k")
	g.Log().Info(ctx, v2.String())

	// setex
	g.Redis().SetEX(ctx, "keyEx", "v4", 2000)
	v3, _ := g.Redis().Get(ctx, "keyEx")
	g.Log().Info(ctx, v3.String())

	// list
	g.Redis().RPush(ctx, "keyList", "v4")
	v4, _ := g.Redis().LPop(ctx, "keyList")
	g.Log().Info(ctx, v4.String())

	// hash
	g.Redis().HSet(ctx, "keyHash", g.Map{"v1": "v5"})
	v5, _ := g.Redis().HGet(ctx, "keyHash", "v1")
	g.Log().Info(ctx, v5.String())

	// set
	g.Redis().SAdd(ctx, "keySet", "v6")
	v6, _ := g.Redis().SPop(ctx, "keySet")
	g.Log().Info(ctx, v6.String())

	// sort set
	g.Redis().ZAdd(ctx, "keySortSet", &gredis.ZAddOption{}, gredis.ZAddMember{Score: 1, Member: "v7"})
	v7, _ := g.Redis().ZRem(ctx, "keySortSet", "v7")
	g.Log().Info(ctx, v7)

}
