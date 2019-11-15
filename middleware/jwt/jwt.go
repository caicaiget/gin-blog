package jwt

import (
	"gin-blog/pkg/util"
	"github.com/b3log/gulu"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := gulu.Ret.NewResult()
		token, err := c.Cookie("Authorization")
		if err != nil {
			result.Msg = "你还没有登录！"
			result.Code = http.StatusUnauthorized
			c.JSON(http.StatusOK, result)
			c.Abort()
			return
		}
		if claims, err := util.ParseToken(token); err != nil{
			result.Msg = "错误的token请登录！"
			result.Code = http.StatusUnauthorized
			c.JSON(http.StatusOK, result)
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			result.Msg = "登录token过期，请重新登录！"
			result.Code = http.StatusUnauthorized
			c.JSON(http.StatusOK, result)
			c.Abort()
			return
		}else {
			//user := models.GetUserById(claims.UserId)
			err = claims.SetUser(c)
			if err != nil {
				result.Msg = err.Error()
				result.Code = http.StatusUnauthorized
				c.JSON(http.StatusOK, result)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
