package main

import (
	"time"

	"gorm.io/gorm"
)

// User和Order是一对多的关系：一个user有多个order，一个order只能属于一个user
// 下面两个struct展示了has-many的关系
// User has many Orders

// OrderHasMany属于User，其中UserID是外键
type OrderHasMany struct {
	gorm.Model
	Num    int    `gorm:"column:num"`
	Item   string `gorm:"column:item"`
	Price  int    `gorm:"column:price"`
	UserID int
}

func (*OrderHasMany) TableName() string {
	return "orders"
}

type UserHasMany struct {
	Id          int       `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	Country     string    `gorm:"column:country"`
	PhoneNumber string    `gorm:"column:phone_number"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	// foreighKey指定了child表（也就是上面的Order）的外键列的名称；references指定了parent表（也就是当前这个User）的被引用的列的名称
	Orders []OrderHasMany `gorm:"foreignKey:UserID;references:Id"`
}

func (*UserHasMany) TableName() string {
	return "users"
}
