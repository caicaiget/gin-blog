package jwt

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token, err := c.Cookie("authorization")
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		}
		if claims, err := util.ParseToken(token); err != nil{
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		}else {
			user := models.GetUserById(claims.UserId)
			c.Set("user", user)
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
