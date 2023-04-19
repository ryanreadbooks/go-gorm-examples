# 让GORM也能识别自定义数据类型

## 实现方式
1. 自定义一个struct
2. 为这个struct实现几个接口，下面有两种可选方案，只要实现其中一种即可
   * database/sql.Scanner接口 + database/sql/driver.Valuer接口 + gorm.GormDataTypeInterface接口
   * database/sql.Scanner接口 + gorm/GormValuerInterface接口 + gorm/GormDataTypeInterface接口