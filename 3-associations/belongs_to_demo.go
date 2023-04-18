package main

import (
	"fmt"

	"gorm.io/gorm"
)

func BelongsToDemo(db *gorm.DB) {
	var orders []OrderBelongsTo
	// preload
	// 在查询orders的时候首先使用preload查出对应的user信息
	// 此处我们同样使用自定义Preload查询语句，从'users'表中执行preload操作
	// 此处的Preload的第一个参数表示的是OrderBelongsTo结构体中的User字段，表示要通过Preload填充这个字段
	res := db.Debug().Table("orders").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Table("users")
	}).Find(&orders)
	if res.Error != nil {
		fmt.Printf("query orders error: %v\n", res.Error)
		return
	}
	for _, o := range orders {
		fmt.Printf("%#v\n", o)
	}
}
