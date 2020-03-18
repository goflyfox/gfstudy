package test

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"testing"
)

var path = "http://127.0.0.1"

func TestGet(t *testing.T) {
	if response, err := ghttp.Get(path); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
	}
}

func TestHello(t *testing.T) {
	if response, err := ghttp.Get(path + "/hello"); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
	}
}

func TestPost(t *testing.T) {
	if response, err := ghttp.Post(path+"/test", "name=john&age=18"); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
	}
}

func TestPostJson(t *testing.T) {
	if response, err := ghttp.Post(path+"/test2",
		`{"passport":"john","password":"123456"}`); err != nil {
		panic(err)
	} else {
		defer response.Close()
		t.Log(response.ReadAllString())
	}
}

func TestPostHeader(t *testing.T) {
	c := ghttp.NewClient()
	c.SetHeader("Cookie", "name=john; score=100")
	if r, e := c.Post(path + "/test3"); e != nil {
		panic(e)
	} else {
		fmt.Println(r.ReadAllString())
	}
}
