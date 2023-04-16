package middleware

import (
	"net/http"
	"strings"

	"blog_server/common"
	"blog_server/model"

	"github.com/gin-gonic/gin"
)

/* middl1e/AuthMiddleware.go */
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.Request.Header.Get("Authorization")
		// Authorization: Bearer xxxxxxxxxxxxxxxxxxx

		// token为空
		if tokenString == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "token参数不能为空", // 权限不足
			})
			c.Abort()
			return
		}
		// 非法token
		if tokenString == "" || len(tokenString) < 7 || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token参数不正确", // 权限不足
			})
			c.Abort()
			return
		}
		// 提取token的有效部分
		tokenString = tokenString[7:]

		// 解析token
		token, claims, err := common.ParseToken(tokenString)

		// token解析失败
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token 认证失败", // 权限不足
			})
			c.Abort()
			return
		}

		// 获取claims中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.Where("id =?", userId).First(&user)

		// 将用户信息写入上下文便于读取
		c.Set("user", user)

		// 中间件拦截器: 请求路由之前的处理
		c.Next()
		// 进行下一步路由的请求

	}
}
