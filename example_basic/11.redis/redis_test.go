package main

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestString(t *testing.T) {
	ctx := gctx.New()
	_, err := g.Redis().Set(ctx, "keyString", "fox")
	if err != nil {
		panic(err)
	}
	v, _ := g.Redis().Get(ctx, "keyString")
	g.Log().Info(ctx, v.String())
	_, _ = g.Redis().Del(ctx, "keyString")

	_ = g.Redis().SetEX(ctx, "keyEx", "fox2", 2000)
	v2, _ := g.Redis().Get(ctx, "keyEx")
	g.Log().Info(ctx, v2.String())
	_, _ = g.Redis().Del(ctx, "keyEx")
}

func TestList(t *testing.T) {
	ctx := gctx.New()
	_, _ = g.Redis().RPush(ctx, "keyList", "v3")
	_, _ = g.Redis().RPush(ctx, "keyList", "v4")
	_, _ = g.Redis().RPush(ctx, "keyList", "v4")
	count, _ := g.Redis().LLen(ctx, "keyList")
	g.Log().Info(ctx, "count:", count)
	val, _ := g.Redis().LPop(ctx, "keyList")
	g.Log().Info(ctx, val.String())
	_, _ = g.Redis().Del(ctx, "keyList")
}

func TestSet(t *testing.T) {
	ctx := gctx.New()
	_, _ = g.Redis().SAdd(ctx, "keySet", "v6")
	_, _ = g.Redis().SAdd(ctx, "keySet", "v7")
	_, _ = g.Redis().SAdd(ctx, "keySet", "v7")
	count, _ := g.Redis().SCard(ctx, "keySet")
	g.Log().Info(ctx, "count:", count)
	val, _ := g.Redis().SPop(ctx, "keySet")
	g.Log().Info(ctx, val.String())
	_, _ = g.Redis().Del(ctx, "keySet")
}

func TestZSet(t *testing.T) {
	ctx := gctx.New()
	_, _ = g.Redis().ZAdd(ctx, "keySortSet", &gredis.ZAddOption{}, gredis.ZAddMember{Score: 1, Member: "v1"})
	_, _ = g.Redis().ZAdd(ctx, "keySortSet", &gredis.ZAddOption{}, gredis.ZAddMember{Score: 3, Member: "v3"})
	_, _ = g.Redis().ZAdd(ctx, "keySortSet", &gredis.ZAddOption{}, gredis.ZAddMember{Score: 2, Member: "v2"})
	vals, _ := g.Redis().ZRange(ctx, "keySortSet", 1, 3)
	g.Log().Info(ctx, "vals:", vals.Strings())
	_, _ = g.Redis().Del(ctx, "keySortSet")
}

func TestHash(t *testing.T) {
	ctx := gctx.New()
	_, _ = g.Redis().HSet(ctx, "keyHash", g.Map{"id": "555", "name": "fox"})
	val, _ := g.Redis().HGet(ctx, "keyHash", "id")
	g.Log().Info(ctx, val.String())
	_, _ = g.Redis().Del(ctx, "keyHash")
}

func TestCache(t *testing.T) {
	ctx := gctx.New()
	redisCache := "cache"
	// 默认default源
	_, err := g.Redis().Set(ctx, "k", "fox")
	// 获取cache链接,DoVar转换
	val, err := g.Redis(redisCache).Get(ctx, "k")
	if err != nil {
		panic(err)
	}
	g.Log().Info(ctx, val.String())
	// setNX
	flag, _ := g.Redis(redisCache).SetNX(ctx, "k", "v")
	g.Log().Info(ctx, flag)
	val, err = g.Redis(redisCache).Get(ctx, "k")
	if err != nil {
		panic(err)
	}
	g.Log().Info(ctx, val.String())
	// 删除数据
	_, _ = g.Redis().Del(ctx, "k")
	_, _ = g.Redis(redisCache).Del(ctx, "k")
}
