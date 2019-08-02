package api

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


func GetAuth(c *gin.Context) {
	var auth Auth
	_ = c.ShouldBindJSON(&auth)

	valid := validation.Validation{}
	valid.Required(auth.Username, "username").Message("username must be")
	valid.Required(auth.Password, "password").Message("password must be")
	if valid.HasErrors() {
		c.JSON(e.InvalidParams, gin.H{
			"msg": valid.Errors[0].Message,
		})
	}

	isExist, id := models.CheckAuth(auth.Username, auth.Password)
	if !isExist {
		c.JSON(e.ErrorAuth, gin.H{
			"msg": "Incorrect account or password",
		})
		return
	}

	token, err := util.GenerateToken(auth.Username, id)
	if err != nil {
		c.JSON(e.ErrorAuth, gin.H{
			"msg": err,
		})
		return
	} else {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:  "Authorization",
			Value: token,
		})
		c.JSON(e.SUCCESS, gin.H{
			"username": auth.Username,
			"id": id,
		})
	}
}


