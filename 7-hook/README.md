# Hook

Hook 是在创建、查询、更新、删除等操作之前、之后调用的函数。

如果您已经为模型定义了指定的方法，它会在创建、更新、查询、删除时自动被调用。如果任何回调返回错误，GORM 将停止后续的操作并回滚事务。

钩子方法的函数签名应该是 func(*gorm.DB) error

## Create Hook触发顺序
```go
// 开始事务
BeforeSave
BeforeCreate
// 关联前的 save
// 插入记录至 db
// 关联后的 save
AfterCreate
AfterSave
// 提交或回滚事务
```

## Update Hook触发顺序
```go
// 开始事务
BeforeSave
BeforeUpdate
// 关联前的 save
// 更新 db
// 关联后的 save
AfterUpdate
AfterSave
// 提交或回滚事务
```

## Delete Hook触发顺序
```go
// 开始事务
BeforeDelete
// 删除 db 中的数据
AfterDelete
// 提交或回滚事务
```

## 查询操作 Hook触发顺序
```go
// 从 db 中加载数据
// Preloading (eager loading)
AfterFind
```
