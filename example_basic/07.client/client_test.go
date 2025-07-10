package main_test

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"net/http"
	"testing"
)

var path = "http://127.0.0.1:8199/api"

// GET/POST请求
func TestWelcome(t *testing.T) {
	ctx := gctx.New()
	if response, err := g.Client().Get(ctx, path); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.RawRequest())
		t.Log(response.RawResponse())
		t.Log(response.ReadAllString())
	}
	if response, err := g.Client().Post(ctx, path); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.Raw()) // 包含请求和响应
		t.Log(response.ReadAllString())
	}
}

// GET请求带参数
func TestGet(t *testing.T) {
	if response, err := g.Client().Get(gctx.New(), path+"/get?name=whoami"); err != nil {
		panic(err)
	} else {
		defer response.Close()
		response.RawDump()
	}
}

// POST请求
func TestPost(t *testing.T) {
	if response, err := g.Client().Post(gctx.New(), path+"/post", "name=fox&age=18"); err != nil {
		panic(err)
	} else {
		defer response.Close()
		response.RawDump()
	}
}

// POST JSON
func TestPostJson(t *testing.T) {
	if response, err := g.Client().Post(gctx.New(), path+"/post/json",
		g.Map{"passport": "fox", "password": "fox123"}); err != nil {
		panic(err)
	} else {
		defer response.Close()
		response.RawDump()
	}
}

// POST Header头
func TestPostHeader(t *testing.T) {
	c := g.Client()
	c.SetHeader("Cookie", "name=fox; score=100")
	//c.SetCookieMap(g.MapStrStr{
	//	"name":  "fox",
	//	"score": "100",
	//})
	if r, e := c.Post(gctx.New(), path+"/post/header"); e != nil {
		panic(e)
	} else {
		defer r.Close()
		r.RawDump()
	}
}

// POST Header头
func TestPostHeaderRaw(t *testing.T) {
	c := g.Client()
	c.SetHeaderRaw(`
accept-encoding: gzip, deflate, br
accept-language: zh-CN,zh;q=0.9,en;q=0.8
referer: https://idonottell.you
cookie: name=fox; score=100
user-agent: my test http client
 	`)
	if r, e := c.Post(gctx.New(), path+"/post/header/raw"); e != nil {
		panic(e)
	} else {
		defer r.Close()
		r.RawDump()
	}
}

// Test Handler
func TestHandler(t *testing.T) {
	c := g.Client()
	c.Use(func(c *gclient.Client, r *http.Request) (resp *gclient.Response, err error) {
		t.Log("Handler前")
		resp, err = c.Next(r)
		t.Log("Handler后")
		return resp, err
	})
	if r, e := c.Post(gctx.New(), path); e != nil {
		panic(e)
	} else {
		defer r.Close()
		r.RawDump()
	}
}
