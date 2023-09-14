package store

import (
    "sync"

    "gorm.io/gorm"
)

var (
    once sync.Once
    // 全局变量，方便其它包直接调用已初始化好的 S 实例.
    S *dataStoreImpl
)

// IStore 定义了 Store 层需要实现的方法.
type IStore interface {
    Users() UserStore
}

// dataStoreImpl 是 IStore 的一个具体实现.
type dataStoreImpl struct {
    db *gorm.DB
}

// 确保 dataStoreImpl 实现了 IStore 接口.
var _ IStore = (*dataStoreImpl)(nil)

// NewStore 创建一个 IStore 类型的实例.
func NewStore(db *gorm.DB) *dataStoreImpl {
    // 确保 S 只被初始化一次
    once.Do(func() {
        S = &dataStoreImpl{db}
    })

    return S
}

// Users 返回一个实现了 UserStore 接口的实例.
func (ds *dataStoreImpl) Users() UserStore {
    return newUsers(ds.db)
}
