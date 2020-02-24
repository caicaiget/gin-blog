package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/b3log/gulu"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

func GetArticles(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)
	userName := c.Param("user")
	user, ok := models.GetUserByName(userName)
	if !ok {
		result.Code = http.StatusNotFound
		result.Msg = "user not found"
		return
	}
	data := make(map[string]interface{})
	pageNum, pageSize := util.GetPage(c)
	var err error
	data["lists"], err = models.GetArticles(pageNum, pageSize, user.ID)
	if err != nil {
		log.Println(err)
	}
	data["count"], err = models.GetArticleCount(user.ID)
	if err != nil {
		log.Println(err)
	}
	result.Data = data
	result.Code = http.StatusOK
}

func CreateArticle(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)
	user := util.GetUser(c)
	var articleJson models.Article
	_ = c.ShouldBindJSON(&articleJson)
	articleJson.CreatedBy = user.UserId
	articleJson.ModifiedBy = user.UserId

	if err := util.ValidStructure(articleJson, c); err != nil {
		result.Msg = err.Error()
		result.Code = http.StatusBadRequest
		return
	}
	if article, err := models.CreateArticle(&articleJson); err != nil {
		result.Msg = err.Error()
		result.Code = http.StatusBadRequest
	} else {
		result.Data = article
		result.Code = http.StatusOK
	}
}

func UpdateArticle(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)
	var article models.Article
	articleId := com.StrTo(c.Param("id")).MustInt64()
	if dbArticle, ok := models.GetArticleById(articleId); !ok {
		result.Msg = "article 在数据库不存在"
		return
	} else {
		article = dbArticle
	}
	_ = c.ShouldBindJSON(&article)
	if err := util.ValidStructure(article, c); err != nil {
		result.Msg = err.Error()
		result.Code = http.StatusBadRequest
		return
	}
	if article, err := models.EditArticle(&article); err != nil {
		result.Msg = err.Error()
		result.Code = http.StatusBadRequest
	} else {
		result.Data = article
		result.Code = http.StatusOK
	}
}

func DeleteArticle(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)
	id := com.StrTo(c.Param("id")).MustInt64()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("id必须大于0")

	if valid.HasErrors() {
		result.Msg = valid.Errors[0].Message
		result.Code = http.StatusBadRequest
		return
	}
	if _, ok := models.GetArticleById(id); ok {
		models.DeleteArticle(id)
		result.Msg = "success"
		result.Code = http.StatusOK
	} else {
		result.Msg = "article 不存在"
		result.Code = http.StatusBadRequest
	}
}
