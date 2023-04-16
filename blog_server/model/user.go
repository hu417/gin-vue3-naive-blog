package model

import "github.com/jinzhu/gorm"

/* model/user.go */
// 用户结构体User
type User struct {
	gorm.Model
	UserName    string `json:"username" gorm:"varchar(20);not null" description:"用户名"`
	PhoneNumber string `json:"phonenumber" gorm:"varchar(20);not null;unique" description:"手机号"`
	Password    string `json:"password" gorm:"size:255;not null" description:"密码"`
	Avatar      string `json:"avatar" gorm:"size:255;not null" description:"头像"`
	Collects    Array  `gorm:"type:longtext" description:"收藏,数据类型为数组"`
	Following   Array  `gorm:"type:longtext" description:"关注,数据类型为数组"`
	Fans        int    `gorm:"AUTO_INCREMENT" description:"粉丝"`
}

// UserInfo为部分的用户信息，便于将数据库的查询结果绑定到结构体上
type UserInfo struct {
	ID       uint   `json:"id"`
	Avatar   string `json:"avatar" description:"头像"`
	UserName string `json:"userName" description:"用户名"`
}
