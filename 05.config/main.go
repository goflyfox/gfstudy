package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
)

func main() {
	// /home/www/templates/
	fmt.Println(g.Cfg().Get("viewpath"))

	// 127.0.0.1:6379,1
	c := g.Cfg()
	fmt.Println(c.Get("redis.cache"))

	// test2
	fmt.Println(c.Get("database.default.1.name"))
}
