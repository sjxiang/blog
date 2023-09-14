package model

import (
	"time"
	"gorm.io/gorm"

	"github.com/sjxiang/blog/pkg/util"
)

// UserM 是数据库中 user 记录 struct 格式的映射.
type UserM struct {
	ID        int64     `gorm:"column:id;primary_key"` 
	Username  string    `gorm:"column:username"`       
	Password  string    `gorm:"column:password"`       
	Nickname  string    `gorm:"column:nickname"`       
	Email     string    `gorm:"column:email"`          
	Phone     string    `gorm:"column:phone"`          
	CreatedAt time.Time `gorm:"column:createdAt"`      
	UpdatedAt time.Time `gorm:"column:updatedAt"`      
}

// TableName 用来指定映射的 MySQL 表名.
func (u *UserM) TableName() string {
	return "user"
}

// BeforeCreate 在创建数据库记录之前加密明文密码.
func (u *UserM) BeforeCreate(tx *gorm.DB) (err error) {
    // Encrypt the user password.
    u.Password, err = util.Encrypt(u.Password)
    if err != nil {
        return err
    }

    return nil
}



// db2struct --gorm --no-json -H 172.21.0.3 -d blog -t user --package model --struct UserM -u root -p '123456' --target=user.go
