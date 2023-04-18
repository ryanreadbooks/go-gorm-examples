package main

import (
	"fmt"

	"gorm.io/gorm"
)

func HasManyDemo(db *gorm.DB) {
	var users []UserHasMany
	// res := db.Debug().Table("users").Preload("Orders").Find(&users)
	// 可以自定义Preload的查询语句
	// 下面的查询的功能：查找users的时候预加载相关的orders
	// Preload的参数解释：
	// 1. 第一个参数："Orders"知名了UserHasMany结构体中，需要预加载的字段的名称
	// 2. 第二个参数：下面我就自定义了执行Preload查询的表格名字，表示从"orders"表中查询数据
	//    Preload的第二个参数是一个可变长的interface{}，所以可以传入函数对象
	res := db.Debug().Table("users").Preload("Orders", func (db *gorm.DB) *gorm.DB {
		return db.Table("orders")
	}).Find(&users)

	if res.Error != nil {
		fmt.Printf("query error: %v", res.Error)
		return
	}
	for _, u := range users {
		fmt.Printf("%#v\n", u)
	}
}

// 插入has-many关系的记录
func HasManyInsertion(db *gorm.DB) {
	u := UserHasMany{
		Name: "Wanda",
		Email: "wanda@foxmail.com",
		Country: "USA",
		PhoneNumber: "17090987654",
		Orders: []OrderHasMany{
			{Num: 10, Item: "cup", Price: 18},
			{Num: 50, Item: "tissue", Price: 129},
			{Num: 1, Item: "tea", Price: 78},
		},
	}
	db.Debug().Create(&u) // 在插入user的同时，也把订单order也插入了
}