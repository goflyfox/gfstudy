package test

import (
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/text/gstr"
	"testing"
)

var path = "http://127.0.0.1/api"

// gstr 示例
func TestStr(t *testing.T) {
	p := fmt.Println
	p("Contains:  ", gstr.Contains("test", "es"))
	p("Count:     ", gstr.Count("test", "t"))
	p("HasPrefix: ", gstr.HasPrefix("test", "te"))
	p("HasSuffix: ", gstr.HasSuffix("test", "st"))
	p("Join:      ", gstr.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", gstr.Repeat("a", 5))
	p("Replace:   ", gstr.Replace("foo", "o", "0", -1))
	p("Replace:   ", gstr.Replace("foo", "o", "0", 1))
	p("Split:     ", gstr.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", gstr.ToLower("TEST"))
	p("ToUpper:   ", gstr.ToUpper("test"))
	p("Trim:   ", gstr.Trim("  test  "))
}

func TestMap(t *testing.T) {
	// 常规map方法
	p := fmt.Println
	p(g.Map{"a": 1, "b": 2})

	// gmap用法
	hash := gmap.New()
	hash.Set("a", 1)
	hash.Sets(map[interface{}]interface{}{
		"a": 1,
		"b": 2,
	})
	p(hash)
}
