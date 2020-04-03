package test

import (
	"github.com/gogf/gf/frame/g"
	"testing"
)

func TestDB(t *testing.T) {
	g.DB().Exec("select 1")
}

func TestRedis(t *testing.T) {
	g.Redis().Do("set", "goframe", "no.1")
	g.Redis().Do("del", "goframe")
}
