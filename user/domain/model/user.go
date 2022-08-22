package model

type User struct {
	// 主键
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	// 用户名称
	UserName string `gorm:"unique_index;not_null"`
	// 添加需要的字段
	FirstName string
	// ...
	// 密码
	HashPassword string
}