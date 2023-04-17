# CRUD接口

## 学习到的方法

### 检索结果

```go
First() 对结果按照主键进行升序排列然后用limit 1返回一条结果
```

```go
Last() 对结果按照主键进行降序排列然后用limit 1返回一条结果
```

```go
Find() 返回所有满足查询条件的结果，这个方法可以给slice赋值
```

```go
Table() 可以指定在哪张表上执行操作
```

### 条件查询

```go
Where() 指定查询结果，多个Where()可以链式调用从而达到条件的拼接(AND操作)。Where()的参数可以用LIKE, IN, != 等语法。
```

```go
Not() NOT操作，即WHERE NOT
```


```go
Select() 选定指定的字段
```

### 排序

```go
Order() 在这个方法里面指定按照哪个字段进行排序，并且可以指定升序还是降序排列
```

### 分页查询

```go
Limit() 表示取多少条记录
```

```go
Offset() 表示偏移多少条记录
```

### 插入

```go
Create()
```

### 更新

```go
Update() 用来更新一个字段，从方法的命名来看，这个方法没有s，所以只是更新一个字段
```

```go
Updates() 用来更新多个字段, 可以通过struct或者map[string]interface{}来指定更新的值，这个方法的名字最后包含一个s，所以可以用来更新多个字段
```

```go
Save()
```

### 删除

gorm还支持软删除，[详见官方资料](https://gorm.io/zh_CN/docs/delete.html)。
```go
Delete() 可以通过指定主键删除特定记录
```