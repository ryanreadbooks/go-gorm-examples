package main

import (
	"time"

	"gorm.io/gorm"
)

func SimpleUpdateDemo(db *gorm.DB) {

	// 使用Table指定要更新哪张表
	db.Debug().Table("users").Where("id = ?", 2).Update("phone_number", "16123123123")

	// 或者用Model来指定更新哪张表
	db.Debug().Model(&User{}).Where("id = ?", 3).Update("phone_number", "17777777777")

	// 更新一行的多列数据
	// 注意这个接口用的是Updates，多了一个s
	// 如果Updates里面用的是struct的话，那么默认只会更新非零值
	db.Debug().Model(&User{}).Where("id = ?", 4).Updates(User{PhoneNumber: "12213213123", Country: "New Zealand", UpdatedAt: time.Now()})
	// 使用map[string]interface{}类型同样可以指定更新的参数，并且可以接受零值
	db.Debug().Model(&User{}).Where("id = ?", 6).Updates(map[string]interface{}{"PhoneNumber": "13613613613", "Country": "Canada", "UpdatedAt": time.Now()})

	// 或者使用Save方法也可以更新记录
	var updatedUser User
	db.Debug().Where(2).First(&updatedUser)
	updatedUser.Email = "jane.smith@foxmail.com" // 只更新字段
	// updatedUser.UpdatedAt = time.Now() // 即使手动指定更新时间，gorm也会帮我们自动更新更新时间，因为gorm约定了UpdatedAt为更新时间，在更新是自动填充时间
	db.Debug().Save(&updatedUser) // 然后更新
}

func UpdateDemo(db *gorm.DB) {
	SimpleUpdateDemo(db)
}
