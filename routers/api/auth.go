package api

import (
	"gin-blog/models"
	"gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/b3log/gulu"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetAuth(c *gin.Context) {
	result := gulu.Ret.NewResult()
	result.Code = http.StatusOK
	defer c.JSON(http.StatusOK, result)
	var auth Auth
	_ = c.ShouldBindJSON(&auth)

	valid := validation.Validation{}
	valid.Required(auth.Username, "username").Message("username must be")
	valid.Required(auth.Password, "password").Message("password must be")
	if valid.HasErrors() {
		result.Code = http.StatusNotFound
		result.Msg = valid.Errors[0].Message
		return
	}

	isExist, id := models.CheckAuth(auth.Username, auth.Password)
	if !isExist {
		result.Code = http.StatusNotFound
		result.Msg = "Incorrect account or password"
		return
	}

	token, err := util.GenerateToken(auth.Username, id)
	if err != nil {
		result.Code = http.StatusNotFound
		result.Msg = err.Error()
	} else {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:  "Authorization",
			Value: token,
		})
		result.Data = gin.H{
			"username": auth.Username,
			"id":       id,
		}
	}
}

func GetLoginUser(c *gin.Context) {

}
