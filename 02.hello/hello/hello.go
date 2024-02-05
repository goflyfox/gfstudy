package main

import (
	"fmt"
	"github.com/gogf/gf"
	"github.com/gogf/gf/crypto/gmd5"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println(gf.VERSION)
	fmt.Println(gmd5.EncryptString("123456"))
}
