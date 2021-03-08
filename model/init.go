package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库连接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string) {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("连接数据库出现异常: %v", err))
	}

	// 设置连接池
	sqlDB, _ := db.DB()
	// 连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(80)
	// 打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(120)
	// 连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	DB = db

	migration()
}
