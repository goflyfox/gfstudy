package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gvalid"
	"testing"
)

// 单条校验
func TestCheck(t *testing.T) {
	rule := "length:6,16"
	if m := gvalid.Check("12345", rule, nil); m != nil {
		t.Log(m)
	} else {
		t.Log("check ok!")
	}
}

// map校验
func TestCheckMap(t *testing.T) {
	params := map[string]interface{}{
		"passport":  "john",
		"password":  "123456",
		"password2": "1234567",
	}
	rules := map[string]string{
		"passport":  "required|length:6,16",
		"password":  "required|length:6,16|same:password2",
		"password2": "required|length:6,16",
	}
	msgs := map[string]interface{}{
		"passport": "账号不能为空|账号长度应当在:min到:max之间",
		"password": map[string]string{
			"required": "密码不能为空",
			"same":     "两次密码输入不相等",
		},
	}
	if e := gvalid.CheckMap(params, rules, msgs); e != nil {
		fmt.Println("#############")
		g.Dump(e.FirstItem())
		fmt.Println("#############")
		g.Dump(e.FirstRule())
		fmt.Println("#############")
		g.Dump(e.FirstString())
		fmt.Println("#############")
		g.Dump(e.Map())
		fmt.Println("#############")
		g.Dump(e.Maps())
		fmt.Println("#############")
		g.Dump(e.String())
		fmt.Println("#############")
		g.Dump(e.Strings())
	} else {
		t.Log("check ok!")
	}
}

// 对象校验
func TestCheckStruct(t *testing.T) {
	type User struct {
		Uid  int    `gvalid:"uid      @integer|min:1#用户UID不能为空"`
		Name string `gvalid:"name     @required|length:6,30#请输入用户名称|用户名称长度非法"`
	}

	user := &User{
		Name: "john",
	}

	// 使用结构体定义的校验规则和错误提示进行校验
	g.Dump(gvalid.CheckStruct(user, nil).Map())

}
