package model

import (
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func GetDbConn(connString string, maxIdleConns, maxOpenConns, maxLifeTime int) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(maxIdleConns)
	//打开
	db.DB().SetMaxOpenConns(maxOpenConns)
	//超时
	db.DB().SetConnMaxLifetime(time.Second *time.Duration(maxLifeTime))

	DB = db

	migration()
}
