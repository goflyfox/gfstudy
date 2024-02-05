package main

import (
	"fmt"
	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/crypto/gmd5"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println(gf.VERSION)
	fmt.Println(gmd5.EncryptString("123456"))
}
