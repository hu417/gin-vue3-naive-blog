package controller

import (
	"strconv"

	"blog_server/common"
	"blog_server/model"
	"blog_server/response"

	"github.com/gin-gonic/gin"
)

/* controller/CollectsController.go */
// Collects 查询收藏
func Collects(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的id
	id := c.Params.ByName("id")
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	// 判断是否已收藏
	for i := 0; i < len(curUser.Collects); i++ {
		if curUser.Collects[i] == id {
			response.Success(c, gin.H{"collected": true, "index": i}, "查询成功")
			return
		}
	}
	response.Success(c, gin.H{"collected": false}, "查询成功")
}

// NewCollect 新增收藏
func NewCollect(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的id
	id := c.Params.ByName("id")
	// 查找用户
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	var newCollects []string
	newCollects = append(curUser.Collects, id)
	// 更新收藏夹
	if err := db.Model(&curUser).Update("collects", newCollects).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	response.Success(c, nil, "更新成功")
}

// UnCollect 取消收藏
func UnCollect(c *gin.Context) {
	db := common.GetDB()
	// 获取用户ID
	user, _ := c.Get("user")
	// 获取path中的index
	index, _ := strconv.Atoi(c.Params.ByName("index"))
	// 查找用户
	var curUser model.User
	db.Where("id = ?", user.(model.User).ID).First(&curUser)
	var newCollects []string
	newCollects = append(curUser.Collects[:index], curUser.Collects[index+1:]...)
	// 更新收藏夹
	if err := db.Model(&curUser).Update("collects", newCollects).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	response.Success(c, nil, "更新成功")
}
