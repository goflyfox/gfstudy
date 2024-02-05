package test

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

var path = "http://127.0.0.1:8199/api"

// GET请求
func TestGet(t *testing.T) {
	ctx := gctx.New()
	if response, err := g.Client().Get(ctx, path); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
	}
	if response, err := g.Client().Post(ctx, path); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
	}
}

// GET请求带参数
func TestHello(t *testing.T) {
	if response, err := g.Client().Get(gctx.New(), path+"/hello?name=whoami"); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
	}
}

// POST请求
func TestPost(t *testing.T) {
	if response, err := g.Client().Post(gctx.New(), path+"/test", "name=john&age=18"); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
	}
}

// POST JSON
func TestPostJson(t *testing.T) {
	if response, err := g.Client().Post(gctx.New(), path+"/test2",
		`{"passport":"john","password":"123456"}`); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
	}
}

// POST Header头
func TestPostHeader(t *testing.T) {
	c := g.Client()
	c.SetHeader("Cookie", "name=john; score=100")
	if r, e := c.Post(gctx.New(), path+"/test3"); e != nil {
		panic(e)
	} else {
		fmt.Println(r.ReadAllString())
	}
}

// POST Header头
func TestPostHeader2(t *testing.T) {
	c := g.Client()
	c.SetHeaderRaw(`
accept-encoding: gzip, deflate, br
accept-language: zh-CN,zh;q=0.9,en;q=0.8
referer: https://idonottell.you
cookie: name=john; score=100
user-agent: my test http client
 	`)
	if r, e := c.Post(gctx.New(), path+"/test4"); e != nil {
		panic(e)
	} else {
		fmt.Println(r.ReadAllString())
	}
}
