package db

import (
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

// MySQLOptions 定义 MySQL 数据库的选项.
type MySQLOptions struct {
	DataSource            string
    MaxIdleConnections    int
    MaxOpenConnections    int
    MaxConnectionLifeTime time.Duration
    LogLevel              int
}



// NewMySQL 使用给定的选项创建一个新的 gorm 数据库实例.
func NewMySQL(opts *MySQLOptions) (*gorm.DB, error) {
    logLevel := logger.Silent
    if opts.LogLevel != 0 {
        logLevel = logger.LogLevel(opts.LogLevel)
    }
    db, err := gorm.Open(mysql.Open(opts.DataSource), &gorm.Config{
        Logger: logger.Default.LogMode(logLevel),
    })
    if err != nil {
        return nil, err
    }

    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }

    // SetMaxOpenConns 设置到数据库的最大打开连接数
    sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

    // SetConnMaxLifetime 设置连接可重用的最长时间
    sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

    // SetMaxIdleConns 设置空闲连接池的最大连接数
    sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

    return db, nil
}
