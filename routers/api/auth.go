package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
)

type Auth struct {
	Username string `form:"username" valid:"Required; MaxSize(50)"`
	Password string `form:"password" valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	var auth Auth
	err := c.ShouldBindJSON(&auth)

	code := e.InvalidParams
	if err == nil {
		isExist := models.CheckAuth(auth.Username, auth.Password)
		if isExist {
			token, err := util.GenerateToken(auth.Username, auth.Password)
			if err != nil {
				code = e.ErrorAuth
			} else {
				code = e.SUCCESS
				http.SetCookie(c.Writer, &http.Cookie {
					Name: "authorization",
					Value:token,
				})
			}

		} else {
			code = e.ErrorAuth
		}
	} else {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": auth,
	})
}
