package test

import (
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"testing"
)

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
	p := fmt.Println
	// 常规map方法
	// 初始化
	m2 := g.Map{"a": 1, "b": 2}
	p(m2)
	// 设置
	m2["c"] = 25
	p(m2)
	// 获取
	p(m2["b"])
	// 删除
	delete(m2, "c")
	// 遍历
	for k, v := range m2 {
		p(k, v)
	}

	p("###########################")

	// 创建一个默认的gmap对象，
	// 默认情况下该gmap对象不支持并发安全特性，
	// 初始化时可以给定true参数开启并发安全特性。
	m := gmap.New()
	// 设置键值对
	for i := 0; i < 10; i++ {
		m.Set(i, i)
	}
	// 查询大小
	p(m.Size())
	// 批量设置键值对(不同的数据类型对象参数不同)
	m.Sets(map[interface{}]interface{}{
		10: 10,
		11: 11,
	})
	p(m.Size())
	// 查询是否存在
	p(m.Contains(1))
	// 查询键值
	p(m.Get(1))
	// 删除数据项
	m.Remove(9)
	p(m.Size())
	// 批量删除
	m.Removes([]interface{}{10, 11})
	p(m.Size())
	// 当前键名列表(随机排序)
	p(m.Keys())
	// 当前键值列表(随机排序)
	p(m.Values())
	// 查询键名，当键值不存在时，写入给定的默认值
	p(m.GetOrSet(100, 100))
	// 删除键值对，并返回对应的键值
	p(m.Remove(100))
	// 遍历map
	m.Iterator(func(k interface{}, v interface{}) bool {
		fmt.Printf("%v:%v ", k, v)
		return true
	})
	// 清空map
	m.Clear()
	// 判断map是否为空
	p(m.IsEmpty())
}

func TestJson(t *testing.T) {
	p := fmt.Println
	// 创建json
	jsonContent := `{"name":"john", "score":"100"}`
	j := gjson.New(jsonContent)
	p(j.Get("name"))
	p(j.Get("score"))

	// 创建json
	j2 := gjson.New(nil)
	j2.Set("name", "John")
	j2.Set("score", 99.5)
	fmt.Printf(
		"Name: %s, Score: %v\n",
		j2.Get("name").String(),
		j2.Get("score").Float64(),
	)
	p(j2.MustToJsonString())

	// struct转json
	type Me struct {
		Name  string `json:"name"`
		Score int    `json:"score"`
	}
	me := Me{
		Name:  "john",
		Score: 100,
	}
	j3 := gjson.New(me)
	p(j3.Get("name"))
	p(j3.Get("score"))
	// 转换回Struct
	Me2 := new(Me)
	if err := j.Scan(Me2); err != nil {
		panic(err)
	}
	fmt.Printf(`%+v`, Me2)
	p()

	// 格式转换
	p("JSON:")
	p(j3.MustToJsonString())
	p("======================")

	p("XML:")
	p(j3.MustToXmlString("document"))
	p("======================")

	p("YAML:")
	p(j3.MustToYamlString())
	p("======================")

	p("TOML:")
	p(j3.MustToTomlString())
}

func TestMd5(t *testing.T) {
	p := fmt.Println
	// md5加密
	p(gmd5.MustEncrypt("123456"))
}

func TestConv(t *testing.T) {
	i := 123.456
	fmt.Printf("%10s %v\n", "Int:", gconv.Int(i))
	fmt.Printf("%10s %v\n", "Int8:", gconv.Int8(i))
	fmt.Printf("%10s %v\n", "Int16:", gconv.Int16(i))
	fmt.Printf("%10s %v\n", "Int32:", gconv.Int32(i))
	fmt.Printf("%10s %v\n", "Int64:", gconv.Int64(i))
	fmt.Printf("%10s %v\n", "Uint:", gconv.Uint(i))
	fmt.Printf("%10s %v\n", "Uint8:", gconv.Uint8(i))
	fmt.Printf("%10s %v\n", "Uint16:", gconv.Uint16(i))
	fmt.Printf("%10s %v\n", "Uint32:", gconv.Uint32(i))
	fmt.Printf("%10s %v\n", "Uint64:", gconv.Uint64(i))
	fmt.Printf("%10s %v\n", "Float32:", gconv.Float32(i))
	fmt.Printf("%10s %v\n", "Float64:", gconv.Float64(i))
	fmt.Printf("%10s %v\n", "Bool:", gconv.Bool(i))
	fmt.Printf("%10s %v\n", "String:", gconv.String(i))

	fmt.Printf("%10s %v\n", "Bytes:", gconv.Bytes(i))
	fmt.Printf("%10s %v\n", "Strings:", gconv.Strings(i))
	fmt.Printf("%10s %v\n", "Ints:", gconv.Ints(i))
	fmt.Printf("%10s %v\n", "Floats:", gconv.Floats(i))
	fmt.Printf("%10s %v\n", "Interfaces:", gconv.Interfaces(i))

	fmt.Println("##############")
	// struct和map转换
	type User struct {
		Uid  int    `c:"uid"`
		Name string `c:"name"`
	}
	// 对象
	m := gconv.Map(User{
		Uid:  1,
		Name: "john",
	})
	fmt.Println(m)

	fmt.Println("##############")
	user := (*User)(nil)
	err := gconv.Struct(m, &user)
	if err != nil {
		panic(err)
	}
	g.Dump(user)
}
