package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"time"
)

type User struct {
	Id          int       `gorm:"column:id;primaryKey"`
	Name        string    `gorm:"column:name"`
	Email       string    `gorm:"column:email"`
	Country     string    `gorm:"column:country"`
	PhoneNumber string    `gorm:"column:phone_number"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

// 

func (u *User) BeforeSave(db *gorm.DB) error {
	fmt.Printf("BeforeSave %#v\n", u)
	return nil
}
func (u *User) BeforeCreate(db *gorm.DB) error {
	fmt.Printf("BeforeCreate %#v\n", u)
	return nil
}

func (u *User) AfterCreate(db *gorm.DB) error {
	fmt.Printf("AfterCreate %#v\n", u)
	return nil
}

func (u *User) AfterSave(db *gorm.DB) error {
	fmt.Printf("AfterSave %#v\n", u)
	return nil
}

type Order struct {
	gorm.Model
	Num    int
	Item   int
	Price  int
	Userid int
}

func main() {
	db, err := gorm.Open(mysql.Open("ryan:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().Model(&User{}).Create(&User{Name: "Tom", Email: "tom@qq.com", Country: "China", PhoneNumber: "123456"})

}
