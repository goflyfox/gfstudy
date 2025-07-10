package test

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

// Insert
func TestInsert(t *testing.T) {
	// INSERT INTO `user`(`name`) VALUES('john')
	_, err := g.Model("user").Data(g.Map{"uid": 10000, "name": "john"}).Insert()
	if err != nil {
		panic(err)
	}
}

// Update
func TestUpdate(t *testing.T) {
	// UPDATE `user` SET `name`='john guo' WHERE name='john'
	_, err := g.Model("user").Data("name", "john guo").
		Where("name", "john").Update()
	if err != nil {
		panic(err)
	}
}

// Delete
func TestDelete(t *testing.T) {
	// DELETE FROM `user` WHERE uid=10
	_, err := g.Model("user").Where("uid", 10000).Delete()
	if err != nil {
		panic(err)
	}
}

// Select Where
func TestWhere(t *testing.T) {
	// INSERT INTO `user`(`name`) VALUES('john')
	_, err := g.Model("user").Data(g.Map{"uid": 10001, "name": "john"}).Insert()
	if err != nil {
		panic(err)
	}
	_, err = g.Model("user").Data(g.Map{"uid": 10002, "name": "john2"}).Insert()
	if err != nil {
		panic(err)
	}
	// 数量
	count, err := g.Model("user").Where("uid", 10001).Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("count:", count)
	// 获取单个值
	v, err := g.Model("user").Where("uid", 10001).Fields("name").Value()
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", v.String())
	// 查询对象
	r, err := g.Model("user").Where("uid", 10002).One()
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", r.Map()["name"])
	// 查询对象
	//l, err := g.Model("user").As("t").Where("t.uid > ?", 10000).All()
	// 也可以简写为 select * from user as t where t.uid > 10000
	l, err := g.Model("user").As("t").All("t.uid > ?", 10000)
	if err != nil {
		panic(err)
	}
	for index, value := range l {
		fmt.Println(index, value["uid"], value["name"])
	}
	g.Model("user").Where("uid", 10001).Delete()
	g.Model("user").Where("uid", 10002).Delete()
}
