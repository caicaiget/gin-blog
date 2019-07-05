package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
)

type ArticleJson struct {
	Title string `form:"title"`
	Desc string `form:"desc"`
	Text string `form:"text"`
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
	//user, _ := c.Get("user")
	//userId := reflect.ValueOf(user).FieldByName("ID").Int()
	var articleJson ArticleJson
	err := c.ShouldBindJSON(&articleJson)
	if err != nil {
		log.Println(err)
		c.JSON(e.InvalidParams, gin.H{
			"msg": err.Error(),
		})
	}
}
