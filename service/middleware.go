package service

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware 鉴权中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 默认双Token放在请求头Authorization中，使用Bearer开头，相互以空格隔开
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(200, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}

		// 使用空格分割authorization,其中 parts[0]="Bearer" parts[1]={access token}
		parts := strings.Split(authHeader, " ")

		// 检查Authorization格式
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(200, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}

		// 解析token
		claims, err := ParseToken(parts[1])
		if err != nil {
			panic(err)
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
