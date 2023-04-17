package main

import "gorm.io/gorm"

func SimpleDeletionDemo(db *gorm.DB) {
	var user User = User{
		Name: "TmpUserName",
		Email: "tmp@example.com",
		Country: "TmpCountry",
		PhoneNumber: "123456",
	}
	
	// 插入一条临时数据
	db.Debug().Save(&user)

	// 可以得到插入的主键，然后现在将其删除
	db.Debug().Delete(&user)	// 如果user中的主键id是有值的，那么就按照主键id将记录删除

	// 同样可以附带where条件
	db.Debug().Save(&user)
	db.Debug().Where("name = ?", user.Name).Delete(&user)
}

func DeletionDemo(db *gorm.DB) {
	SimpleDeletionDemo(db)
}