package biz

import (
	"github.com/sjxiang/blog/internal/blog/biz/user"
	"github.com/sjxiang/blog/internal/blog/store"
)

// IBiz 定义了 Biz 层需要实现的方法.
type IBiz interface {
	Users() user.UserBiz
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*bizImpl)(nil)

// biz 是 IBiz 的一个具体实现.
type bizImpl struct {
	ds store.IStore
}

// NewBiz 创建一个 IBiz 类型的实例.
func NewBiz(ds store.IStore) *bizImpl {
	return &bizImpl{ds: ds}
}

// Users 返回一个实现了 UserBiz 接口的实例.
func (b *bizImpl) Users() user.UserBiz {
	return user.New(b.ds)
}