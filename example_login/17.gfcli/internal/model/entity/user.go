// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	Id         int    `json:"id"         ` // 主键
	Uuid       string `json:"uuid"       ` // UUID
	LoginName  string `json:"loginName"  ` // 登录名/11111
	Password   string `json:"password"   ` // 密码
	RealName   string `json:"realName"   ` // 真实姓名
	Enable     int    `json:"enable"     ` // 是否启用//radio/1,启用,2,禁用
	UpdateTime string `json:"updateTime" ` // 更新时间
	UpdateId   int    `json:"updateId"   ` // 更新人
	CreateTime string `json:"createTime" ` // 创建时间
	CreateId   int    `json:"createId"   ` // 创建者
}
