package jwt

import (
	"gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie("Authorization")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg" : "你还没有登录！",
			})
			c.Abort()
			return
		}
		if claims, err := util.ParseToken(token); err != nil{
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg" : "错误的token请登录！",
			})
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg" : "登录token过期，请重新登录！",
			})
			c.Abort()
			return
		}else {
			//user := models.GetUserById(claims.UserId)
			err = claims.SetUser(c)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg" : err,
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
