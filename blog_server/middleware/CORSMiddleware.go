package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* middleware/CORSMiddleware.go/ */
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")         // 允许所有域名访问
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")        // 允许的请求类型
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")        // 允许的请求头字段
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true") // 允许后续请求携带认证信息（cookies）
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")          // 预检结果缓存时间
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
