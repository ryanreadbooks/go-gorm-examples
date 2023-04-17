# 使用GORM连接到MySQL数据库

[官方参考资料](https://gorm.io/zh_CN/docs/connecting_to_the_database.html)

## 导包
要导入两个相关的包
```bash
go get gorm.io/gorm
go get gorm.io/driver/mysql
```
## 连接

### DSN格式
DSN: 这个连接的格式和go-sql-driver/mysql中的格式一致

### 连接的函数
```go
gorm.Open()
```


