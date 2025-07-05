package main

import (
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestMd5(t *testing.T) {
	md5, _ := gmd5.EncryptString("123456")
	fmt.Println(md5)
}

func TestMd5Salt(t *testing.T) {
	md5, _ := gmd5.EncryptString("123456")
	fmt.Println(md5)
	fmt.Println(gmd5.EncryptString(md5 + "123456"))
}

func TestBcrypt(t *testing.T) {
	passwordOK := "123456"
	passwordOK, _ = gmd5.EncryptString(passwordOK)
	passwordERR := "12345678"
	passwordERR, _ = gmd5.EncryptString(passwordERR)

	hash, err := bcrypt.GenerateFromPassword([]byte(passwordOK), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(hash)

	encodePW := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	fmt.Println("###", encodePW)
	hash, err = bcrypt.GenerateFromPassword([]byte(passwordOK), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	encodePW = string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	fmt.Println("###", encodePW)
	// 其中：$是分割符，无意义；2a是bcrypt加密版本号；10是cost的值；而后的前22位是salt值；
	// 再然后的字符串就是密码的密文了。

	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordOK))
	if err != nil {
		fmt.Println("pw wrong")
	} else {
		fmt.Println("pw ok")
	}

	// 错误密码验证
	err = bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(passwordERR))
	if err != nil {
		fmt.Println("pw wrong")
	} else {
		fmt.Println("pw ok")
	}
}
