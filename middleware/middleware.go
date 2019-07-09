package middleware

import (
	"gin-blog/pkg/e"
	"github.com/gin-gonic/gin"
	"strings"
)

func JsonAccept() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			mimeType := c.GetHeader("Content-Type")
			if !(mimeType == "application/json" || (strings.HasPrefix(mimeType, "application/") &&
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
