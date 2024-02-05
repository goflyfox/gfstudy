// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User is the golang structure of table sys_user for DAO operations like Where/Data.
type User struct {
	g.Meta     `orm:"table:sys_user, do:true"`
	Id         interface{} // 主键
	Uuid       interface{} // UUID
	LoginName  interface{} // 登录名/11111
	Password   interface{} // 密码
	RealName   interface{} // 真实姓名
	Enable     interface{} // 是否启用//radio/1,启用,2,禁用
	UpdateTime interface{} // 更新时间
	UpdateId   interface{} // 更新人
	CreateTime interface{} // 创建时间
	CreateId   interface{} // 创建者
}
