package main

import (
	"fmt"
	"github.com/gogf/gf/text/gstr"
)

func main() {
	TestGstr()
	TestGmap()
}

func TestGstr() {
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

func TestGmap() {

}
