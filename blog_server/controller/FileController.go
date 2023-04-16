package controller

import (
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

/* controller/FileController.go */
// Upload 上传图像
func Upload(c *gin.Context) {
	// 绑定请求参数file
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "格式错误",
		})
		return
	}

	// 图片名
	filename := header.Filename
	// 图片格式
	ext := path.Ext(filename)
	// 用上传时间作为文件名
	name := "image_" + time.Now().Format("20060102150405")
	newFilename := name + ext
	// 创建文件
	out, err := os.Create("static/img/" + newFilename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "创建错误",
		})
		return
	}
	defer out.Close()

	// 数据复制
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "复制错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"filePath": "/img/" + newFilename},
		"msg":  "上传成功",
	})
}
