package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql" // mysql驱动 go get gorm.io/driver/mysql
	"gorm.io/gorm"         // go get gorm.io/gorm
)

// 演示怎样使用gorm连接到数据库mysql
func main() {
	dsn := "ryan:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	var db *gorm.DB
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error when open db: %v\n", err)
	}
	fmt.Printf("%#v\n", db)
}
