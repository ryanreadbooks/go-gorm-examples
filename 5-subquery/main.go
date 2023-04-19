package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"fmt"
)

type User struct {
	Id          int       `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	Country     string    `gorm:"column:country"`
	PhoneNumber string    `gorm:"column:phone_number"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

type Order struct {
	gorm.Model
	Num int
	Item int 
	Price int
	Userid int
}

func main() {
	db, err := gorm.Open(mysql.Open("ryan:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// select * from users where id in (select user_id from orders where price > 1000);
	// 比如将上面这条sql翻译成gorm
	var users []User
	db.Debug().Model(&User{}).Where("id in (?)", db.Table("orders").Select("user_id").Where("price > ?", 1000)).Find(&users)
	for _, u := range users {
		fmt.Printf("%#v\n", u)
	}
}
