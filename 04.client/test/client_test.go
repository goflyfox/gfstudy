package test

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"testing"
)

var path = "http://127.0.0.1"

// GET请求
func TestGet(t *testing.T) {
	if response, err := ghttp.Get(path); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
	}
}

// GET请求带参数
func TestHello(t *testing.T) {
	if response, err := ghttp.Get(path + "/hello?name=whoami"); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
	}
}

// POST请求
func TestPost(t *testing.T) {
	if response, err := ghttp.Post(path+"/test", "name=john&age=18"); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
	}
}

// POST JSON
func TestPostJson(t *testing.T) {
	if response, err := ghttp.Post(path+"/test2",
		`{"passport":"john","password":"123456"}`); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
	}
}

// POST Header头
func TestPostHeader(t *testing.T) {
	c := ghttp.NewClient()
	c.SetHeader("Cookie", "name=john; score=100")
	if r, e := c.Post(path + "/test3"); e != nil {
		panic(e)
	} else {
		fmt.Println(r.ReadAllString())
	}
}

// POST Header头
func TestPostHeader2(t *testing.T) {
	c := ghttp.NewClient()
	c.SetHeaderRaw(`
     accept-encoding: gzip, deflate, br
     accept-language: zh-CN,zh;q=0.9,en;q=0.8
     referer: https://idonottell.you
     cookie: name=john; score=100
     user-agent: my test http client
 	`)
	if r, e := c.Post(path + "/test4"); e != nil {
		panic(e)
	} else {
		fmt.Println(r.ReadAllString())
	}
}
