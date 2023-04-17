package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func SimpleInsertion(db *gorm.DB) {
	newUser := &User{Name: "Alen Luis", Email: "alen@google.com", Country: "UK", PhoneNumber: "18888888888", UpdatedAt: time.Now()}

	// 使用Create来插入一条数据
	err := db.Debug().Create(newUser).Error
	if err != nil {
		// 可能出错，插入失败
		fmt.Println(err)
	}
	// 插入成功之后，可以获得插入后的主键
	fmt.Printf("newUser.Id = %d\n", newUser.Id)

	// 也可以一次性插入多条记录
	manyUsers := []User{
		{Name: "San Zhang", Email: "zhangsan@163.com", Country: "China", PhoneNumber: "12212121212", UpdatedAt: time.Now()},
		{Name: "Li S", Email: "lisi123@163.com", Country: "China", PhoneNumber: "18811111111", UpdatedAt: time.Now()},
		{Name: "Huarui Qi", Email: "ruihual@163.com", Country: "China", PhoneNumber: "15555596969", UpdatedAt: time.Now()},
	}

	db.Debug().Create(&manyUsers)
	// 插入成功的话，可以获得主键
	for idx, u := range manyUsers {
		fmt.Printf("manyUsers[%d].Id=%d\n", idx, u.Id)
	}
}

// 同样支持从map中插入数据
func SimpleInertion2(db *gorm.DB) {
	// 一条数据
	// 这个map的key为列名，value为对应列名的值，类型是interface{}
	newUser := map[string]interface{}{
		"Name": "Pablo Álvarez",
		"Email" : "pablo@google.com",
		"Country" : "Spain",
		"PhoneNumber" : "7856215",
		"UpdatedAt" : time.Now(),
	}
	// 然后指定Model并插入
	// 需要注意的是，如果是用map的方式来插入的话，无法得到插入的主键
	db.Debug().Model(&User{}).Create(&newUser)
}

func InsertionDemo(db *gorm.DB) {
	SimpleInsertion(db)
	SimpleInertion2(db)
}
