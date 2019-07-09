package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"reflect"
)

type ArticleJson struct {
	Title string `json:"title" valid:"required,maxsize(16)"`
	Desc string `json:"desc" valid:"maxsize(16)"`
	Text string `json:"text" valid:"required"`
	CreatedBy int
	ModifyBy int
}

func GetArticles(c *gin.Context) {
	user, _ := c.Get("user")
	userId := reflect.ValueOf(user).FieldByName("ID").Interface()
	log.Println()
	data := make(map[string]interface{})
	pageNum, pageSize := util.GetPage(c)
	var err error
	data["lists"], err = models.GetArticles(pageNum, pageSize, userId)
	if err != nil {
		log.Println(err)
	}
	data["count"], err = models.GetArticleCount(userId)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func CreateArticle(c *gin.Context) {
	user, _ := c.Get("user")
	userId := com.StrTo(reflect.ValueOf(user).FieldByName("ID").String()).MustInt()
	var articleJson ArticleJson
	okJson := c.ShouldBindJSON(&articleJson)
	if okJson != nil {
		c.JSON(e.InvalidParams, gin.H{
			"msg": "Content-Type must be json",
		})
	}
	articleJson.CreatedBy = userId
	articleJson.ModifyBy = userId
	err := binding.Validator.ValidateStruct(articleJson)
	if err != nil {
		log.Println(err)
		c.JSON(e.InvalidParams, gin.H{
			"msg": err.Error(),
		})
	}
}
