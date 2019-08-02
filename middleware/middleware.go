package middleware

import (
	"gin-blog/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JsonAccept() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			mimeType := c.GetHeader("Content-Type")
			if !(strings.HasPrefix(mimeType, "application/json") || (strings.HasPrefix(mimeType, "application/") &&
			strings.HasSuffix(mimeType, "+json"))){
			c.JSON(e.InvalidParams, gin.H{
			"msg": "Content-Type must be Json and Json objects are required",
			})
			c.Abort()
			return
			}
		}
		c.Next()
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.GetHeader("Origin")
		if origin == "" {
			//c.JSON(e.InvalidParams, gin.H{
			//	"msg": "Origin is required for headers",
			//})
			//c.Abort()
			//return
			origin = "*"
		}
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
