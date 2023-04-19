package main

import (
	"bytes"
	"context"
	"database/sql/driver"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FullName struct {
	FirstName string
	LastName  string
}

type HomeAddress struct {
	Province string
	City     string
	County   string
}

// 有一个字段用的是自定义的类型FullName
// 为了能够让GORM识别这个自定义类型，我们需要为FullName实现Valuer和Scanner接口
type Person struct {
	gorm.Model
	FullName FullName    `gorm:"column:fullname"`
	Address  HomeAddress `gorm:"column:address"`
}

func (*Person) TableName() string {
	return "persons"
}

// 实现sql.Scanner接口, 实现一个Scan方法
func (f *FullName) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to convert %v to []byte", value)
	}
	// bytes的内容解析后赋值给f
	fmt.Println("bytes=", string(bytes))
	// 这一部分如何实现就是根据每个自定义类型的需求而有所不同
	return nil
}

// 实现sql包中的drive.Valuer接口，实现一个Value方法
func (f FullName) Value() (driver.Value, error) {
	// 将f转换编码成driver.Value格式的内容
	// driver.Value支持以下的数据类型，也就是要将f编码成以下的类型之一：
	// int64 float64 bool string []byte time.Time
	return f.FirstName + " " + f.LastName, nil	// 转换成可以插入数据库中的值
}

// 实现GormDataTypeInterface接口，这个接口要求实现一个方法GormDataType
// 返回自定义类型的数据类型
func (FullName) GormDataType() string {
	return "varchar(255)" // 也就是FullName这个字段以varchar(255)格式类型存储在数据库的表中
}

func (addr *HomeAddress) Scan(value interface{}) error {
	bytesValue, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to convert %v into []byte", value)
	}
	fmt.Println("bytes=", string(bytesValue))

	s := bytes.Split(bytesValue, []byte(", "))
	addr.Province = string(s[0])
	addr.City = string(s[1])
	addr.County = string(s[2])

	return nil
}

// 实现GormValuerInterface接口
func (addr HomeAddress) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	// 用SQL的形式执行value
	return clause.Expr{
		SQL:  "CONCAT(?, ', ', ?, ', ', ?)",
		Vars: []interface{}{addr.Province, addr.City, addr.County},
	}
}

func (HomeAddress) GormDataType() string {
	return "varchar(255)"
}

func main() {

	db, err := gorm.Open(mysql.Open("ryan:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Person{})

	// 插入一条数据
	p := Person{FullName: FullName{FirstName: "San", LastName: "Zhang"}, Address: HomeAddress{Province: "Guangdong", City: "Guangzhou", County: "Tianhe"}}
	db.Debug().Create(&p)

	var p1 Person
	db.First(&p1)
	fmt.Printf("%#v\n", p1)

}
