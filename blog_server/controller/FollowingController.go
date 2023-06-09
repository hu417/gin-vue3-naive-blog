package controller

import (
	"strconv"

	"blog_server/common"
	"blog_server/model"
	"blog_server/response"

	"github.com/gin-gonic/gin"
)

// Following 查询关注
func Following(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的id
	id := c.Params.ByName("id")
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	// 判断是否已关注
	for i := 0; i < len(curUser.Following); i++ {
		if curUser.Following[i] == id {
			response.Success(c, gin.H{"followed": true, "index": i}, "查询成功")
			return
		}
	}
	response.Success(c, gin.H{"followed": false}, "查询成功")
}

// NewFollow 新增关注
func NewFollow(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的id
	id := c.Params.ByName("id")
	// 查找用户
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	//var newFollowing []string
	newFollowing := append(curUser.Following, id)
	// 更新关注列表
	if err := db.Model(&curUser).Update("following", newFollowing).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	// 更新粉丝数
	var followUser model.User
	db.Where("id = ?", id).First(&followUser)
	if err := db.Model(&followUser).Update("fans", followUser.Fans+1).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	response.Success(c, nil, "更新成功")
}

// UnFollow 取消关注
func UnFollow(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的index
	index, _ := strconv.Atoi(c.Params.ByName("index"))
	// 查找用户
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	//var newFollowing []string
	newFollowing := append(curUser.Following[:index], curUser.Following[index+1:]...)
	followId := curUser.Following[index]
	// 更新关注列表
	if err := db.Model(&curUser).Update("following", newFollowing).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	// 更新粉丝数
	var followUser model.User
	db.Where("id = ?", followId).First(&followUser)
	if err := db.Model(&followUser).Update("fans", followUser.Fans-1).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	response.Success(c, nil, "更新成功")
}
