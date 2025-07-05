package test

import (
	"fmt"
	"github.com/gogf/gf/v2/text/gregex"
	"testing"
)

// IsMatch
func TestIsMatch(t *testing.T) {
	// 校验时间是否合法
	var pattern = `\d{4}-\d{2}-\d{2}`
	s1 := []byte(`2019-07-20`)
	fmt.Println("IsMatch1", gregex.IsMatch(pattern, s1))
	pattern = `[21]\d{3}-\d{1,2}-\d{1,2}`
	fmt.Println("IsMatch2", gregex.IsMatch(pattern, s1))
}

// IsMatchString
func TestIsMatchString(t *testing.T) {
	var pattern = `[21]\d{3}-[01]?\d-[0123]?\d`
	s1 := `2019-07-20`
	fmt.Println("IsMatchString", gregex.IsMatchString(pattern, s1))
}

var (
	textStr     = "123 xiangyu liubang xiangyu liubang"
	patternStr  = `\d+\s(\w+)\s\w+\s\w+\s\w+`
	patternStr2 = `\d+\s(\w+)`
	patternStr3 = `(\w+)\sliubang`
)

// Match
func TestMatch(t *testing.T) {
	subs, err := gregex.Match(patternStr, []byte(textStr))
	if err != nil {
		t.Error("Match", err)
	}
	fmt.Println("Match", string(subs[0]), "##group:", string(subs[1]), err)
}

// MatchString
func TestMatchString(t *testing.T) {
	// 匹配全部内容
	subs, err := gregex.MatchString(patternStr, textStr)
	if err != nil {
		t.Error("MatchString", err)
	}
	fmt.Println("MatchString", subs[0], "##group:", subs[1], err)

	// 匹配部分内容
	subs, err = gregex.MatchString(patternStr2, textStr)
	if err != nil {
		t.Error("MatchString2", err)
	}
	fmt.Println("MatchString2", subs[0], "##group:", subs[1], err)
}

// MatchAll
func TestMatchAll(t *testing.T) {
	allGroup, err := gregex.MatchAll(patternStr3, []byte(textStr))
	if err != nil {
		t.Error("MatchAll", err)
	}
	fmt.Println("MatchAll", string(allGroup[0][0]), "##group:", string(allGroup[0][1]), err)
	fmt.Println("MatchAll", string(allGroup[1][0]), "##group:", string(allGroup[1][1]), err)
}

// MatchAllString
func TestMatchAllString(t *testing.T) {
	allGroup, err := gregex.MatchAllString(patternStr3, textStr)
	if err != nil {
		t.Error("MatchAllString", err)
	}
	fmt.Println("MatchAllString", allGroup, "##group:", allGroup[0][1], err)
}

// Replace
func TestReplace(t *testing.T) {
	replace, err := gregex.Replace(patternStr3, []byte("zhuyuanzhang chenyouliang"), []byte(textStr))
	if err != nil {
		t.Error("Replace", err)
	}
	fmt.Println("Replace", string(replace), "##src:", textStr, err)

}

// ReplaceString
func TestReplaceString(t *testing.T) {
	replacedStr, err := gregex.ReplaceString(patternStr3, "zhuyuanzhang chenyouliang", textStr)
	if err != nil {
		t.Error("ReplaceString", err)
	}
	fmt.Println("ReplaceString", replacedStr, "##src:", textStr, err)
}

// Split
func TestSplit(t *testing.T) {
	items := gregex.Split(`\sxiangyu\s`, textStr)
	fmt.Println("Split", items, "###0:", items[0], "##src:", textStr)
}
