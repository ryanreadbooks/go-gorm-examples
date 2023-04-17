package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// 简单的查询的用法
func QueryDemoSimple(db *gorm.DB) {
	// 查出单个对象
	var user User
	// Debug()可以看到执行了怎样的SQL语句
	result := db.Debug().First(&user)
	// 返回的结果的类型是 *gorm.DB (我猜可能是为了方便链式调用)
	// 通过 *gorm.DB取得查询的结果，比如查出了多少条数据，是否有错误等
	// 而查出的行记录则放进了user对象中
	// 如果没有找到记录的话，则出现gorm.ErrRecordNotFound错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("can not find record")
	}
	fmt.Printf("rowAffected = %d, queried user = %#v\n", result.RowsAffected, user)

	// 获取所有的用户，也就是将查询结果放在一个slice里面
	var users []User
	result = db.Debug().Find(&users) // select * from users
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("rowAffected = %d\n", result.RowsAffected)
		for _, user := range users {
			fmt.Printf("%#v\n", user)
		}
	}
}

// 带有条件的查询
// 要用到Where
func QueryDemoCondition(db *gorm.DB) {
	// =
	var user User
	db.Debug().Where("name = ?", "Alice Brown").First(&user) // 只取出满足条件的行记录中的第一行(先按主键排序再limit取第一行)
	fmt.Printf("%#v\n", user)

	// 获取满足条件的所有行记录
	var users []User
	db.Debug().Where("name != ?", "Alice Brown").Find(&users)
	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}

	// IN 操作
	// IN 后面跟的条件用slice的形式指定
	db.Debug().Where("country in ?", []string{"USA", "UK", "China"}).Find(&users)
	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}

	// LIKE 操作
	db.Debug().Where("email like ?", "%@google.com").Find(&users)
	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}

	// AND 多个条件
	// 写法上就是多个Where调用连在一起
	db.Debug().Where("email like ?", "%@example.com").Where("country = ?", "Australia").Find(&users)
	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}
	// 又或者是直接在查询条件中显式写and
	db.Debug().Where("email like ? and country = ?", "%@example.com", "Australia").Find(&users)
	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}
	// 或者使用map来添加指定查询条件
	// 使用map指定的查询条件的值可以为零值
	db.Debug().Where(map[string]string{"name": "Alice Brown", "country": "Australia"}).Find(&users)
	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}

	// NOT
	// 条件里面用int slice的话，表示的是主键的条件
	// 像下面这样的就是查找主键不是1，2，3，4的行记录
	db.Debug().Not([]int{1, 2, 3, 4}).Find(&users)
	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}

	// OR条件
	db.Debug().Where("name = ?", "John Doe").Or("country = ?", "China").Find(&users)
	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}

	// 选定特定的字段
	db.Debug().Select("name", "country").Where("name = ?", "John Doe").Find(&users)
	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}
}

func QueryDemoOrderBy(db *gorm.DB) {

	var users []User

	// ORDER BY 操作
	db.Debug().Where("country in ?", []string{"USA", "UK"}).Order("phone_number desc").Find(&users)
	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}

	db.Debug().Select("updated_at").Order("updated_at desc").Find(&users)
	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}
}

// limit offset
func QueryDemoLimit(db *gorm.DB) {
	// mysql的limit语法： limit [offset,] rows 表示从offset+1(包含)开始，一共去rows行的数据
	var users []User
	db.Debug().Order("id asc").Limit(3).Offset(2).Find(&users)
	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}
}

func QueryDemo(db *gorm.DB) {
	QueryDemoSimple(db)
	fmt.Println("=============================================")
	QueryDemoCondition(db)
	fmt.Println("=============================================")
	QueryDemoOrderBy(db)
	fmt.Println("=============================================")
	QueryDemoLimit(db)
}
