package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func GetArticles(c *gin.Context) {
	user := util.GetUser(c)
	log.Println()
	data := make(map[string]interface{})
	pageNum, pageSize := util.GetPage(c)
	var err error
	data["lists"], err = models.GetArticles(pageNum, pageSize, user.UserId)
	if err != nil {
		log.Println(err)
	}
	data["count"], err = models.GetArticleCount(user.UserId)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, data)
}


func CreateArticle(c *gin.Context) {

	user := util.GetUser(c)

	var articleJson models.Article
	_ = c.ShouldBindJSON(&articleJson)
	articleJson.CreatedBy = user.UserId
	articleJson.ModifiedBy = user.UserId

	if ok := util.ValidStructure(articleJson, c); !ok {
		return
	}

	if article, err := models.CreateArticle(&articleJson); err != nil {
		c.JSON(e.ERROR, gin.H{
			"msg": err,
		})
	} else {
		c.JSON(e.SUCCESS, gin.H{
			"data": article,
		})
	}
}

func UpdateArticle(c *gin.Context) {
	var article models.Article
	articleId := com.StrTo(c.Param("id")).MustInt64()
	if dbArticle, ok := models.GetArticleById(articleId); !ok{
		c.JSON(e.InvalidParams, gin.H{
			"msg": "article 在数据库不存在",
		})
		return
	} else {
		article = dbArticle
	}
	_ = c.ShouldBindJSON(&article)
	if ok := util.ValidStructure(article, c); !ok {
		return
	}
	if article, err := models.EditArticle(&article); err != nil{
		c.JSON(e.ERROR, gin.H{
			"msg": err,
		})
	} else {
		c.JSON(e.SUCCESS, gin.H{
			"data": article,
		})
	}
}

func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于0")

	if valid.HasErrors() {
		c.JSON(e.InvalidParams, gin.H{
			"msg": valid.Errors[0].Message,
		})
		return
	}
	if _, ok := models.GetArticleById(id); ok {
		models.DeleteArticle(id)
		c.JSON(e.SUCCESS, gin.H{
			"msg": "success",
		})
		return
	} else {
		c.JSON(e.InvalidParams, gin.H{
			"msg": "The Article don't exist",
		})
		return
	}
}
