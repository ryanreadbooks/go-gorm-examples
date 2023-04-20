package main

import (
	"errors"
	"fmt"
	"strconv"
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

	// GORM 默认会在执行写入操作（insert，update，delete）的时候开启事务

	// 显式开始事务执行SQL
	db.Transaction(func(tx *gorm.DB) error {
		// 然后再这个函数里面使用tx来执行操作

		u := &User{Name: "Pep", Email: "pep.football@man.com", Country: "UK", PhoneNumber: "23412312312"}
		err := tx.Debug().Model(u).Where("name = ?", u.Name).First(u).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = tx.Debug().Model(u).Create(u).Error
			// 任何的错误发生都会自动回滚事务
			if err != nil {
				fmt.Println(err)
				return err
			}
		}

		u.PhoneNumber = strconv.FormatInt(time.Now().Unix(), 10)
		err = tx.Debug().Model(u).Update("phone_number", u.PhoneNumber).Where(u).Error
		if err != nil {
			fmt.Println(err)
			return err
		}

		// 返回nil提交事务
		return nil
	})

}
