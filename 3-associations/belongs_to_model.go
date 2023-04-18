package main

import (
	"gorm.io/gorm"
	"time"
)

// 下面两个struct的关系是belongs-to关系
// Order结构体belongs to User结构体
// 在表中Order有外键去引用User
type OrderBelongsTo struct {
	gorm.Model        // 包含ID CreatedAt UpdatedAt DeleteAt字段
	Num        int    `gorm:"column:num"`
	Item       string `gorm:"column:item"`
	Price      int    `gorm:"column:price"`
	UserID     int
	// foreignkey指明了现在这个struct中的哪个字段是外键，refereces指定了引用的struct的字段名称
	User UserBelongsTo `gorm:"foreignKey:UserID;references:Id"`
}

type UserBelongsTo struct {
	Id          int       `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	Country     string    `gorm:"column:country"`
	PhoneNumber string    `gorm:"column:phone_number"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

// 可以通过TableName方法指定这个struct对应的是哪一个表格名称
// 实现Tabler接口
func (*OrderBelongsTo) TableName() string {
	return "orders"
}

func (*UserBelongsTo) TableName() string {
	return "users"
}
