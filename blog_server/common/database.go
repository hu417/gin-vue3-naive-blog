package common

import (
	"fmt"
	"net/url"

	"blog_server/model"

	"github.com/jinzhu/gorm"
)

/* common/database.go */
var DB *gorm.DB

// InitDB() 数据库初始化
func InitDB() *gorm.DB {

	driverName := "mysql"
	user := "root"
	password := "123456"
	host := "10.0.0.91"
	port := "3306"
	database := "blog-community"
	charset := "utf8"
	loc := "Asia/Shanghai"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		user,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))

	// 连接数据库
	var err error
	DB, err = gorm.Open(driverName, args)
	if err != nil {
		panic("failed to open database: " + err.Error())
	}

	// 开启日志
	DB.LogMode(true)

	// 迁移数据表
	DB.AutoMigrate(
		&model.User{},
		&model.Category{},
		&model.Article{},
	)

	return DB
}

// 数据库信息获取
func GetDB() *gorm.DB {
	return DB
}
