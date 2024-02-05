package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func main() {
	// redis字符串操作
	g.Redis().Do("SET", "k", "v")
	v, _ := g.Redis().Do("GET", "k")
	g.Log().Info(gconv.String(v))

	// 获取cache链接
	v2, _ := g.Redis("cache").Do("GET", "k")
	g.Log().Info(gconv.String(v2))

	// DoVar转换
	v3, _ := g.Redis().DoVar("GET", "k")
	g.Log().Info(v3.String())

	// setex
	g.Redis().Do("SETEX", "keyEx", 2000, "v4")
	v4, _ := g.Redis().DoVar("GET", "keyEx")
	g.Log().Info(v4.String())

	// list
	g.Redis().Do("RPUSH", "keyList", "v5")
	v5, _ := g.Redis().DoVar("LPOP", "keyList")
	g.Log().Info(v5.String())

	// hash
	g.Redis().Do("HSET", "keyHash", "v1", "v6")
	v6, _ := g.Redis().DoVar("HGET", "keyHash", "v1")
	g.Log().Info(v6.String())

	// set
	g.Redis().Do("SADD", "keySet", "v7")
	v7, _ := g.Redis().DoVar("SPOP", "keySet")
	g.Log().Info(v7.String())

	// sort set
	g.Redis().Do("ZADD", "keySortSet", 1, "v8")
	v8, _ := g.Redis().DoVar("ZREM", "keySortSet", "v8")
	g.Log().Info(v8.Int())

}
