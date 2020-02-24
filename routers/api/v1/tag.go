package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/b3log/gulu"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

type Tag struct {
	Name  string
	State int64
}

//获取多个文章标签
func GetTags(c *gin.Context) {
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
	data["lists"] = models.GetTags(pageNum, pageSize, user.ID)
	data["total"] = models.GetTagTotal(user.ID)

	result.Code = http.StatusOK
	result.Data = data
}

// @Summary Get multiple article tags
// @Produce json
// @Param name query string false "name"
// @Param state query int64 false "state"
// @Param state query int64 false "createdBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":500,"data":{},"msg":"error"}"
// @Router /api/v1/tags [get]
func AddTag(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)
	var tag Tag
	_ = c.ShouldBindJSON(&tag)
	user := util.GetUser(c)

	valid := validation.Validation{}
	valid.Required(tag.Name, "name").Message("名称不能为空")
	valid.MaxSize(tag.Name, 100, "name").Message("名称最长为100字符")
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许0或1")

	if valid.HasErrors() {
		result.Code = http.StatusNotFound
		result.Msg = valid.Errors[0].Message
		return
	}
	if models.ExistTagByName(tag.Name) {
		result.Code = http.StatusNotFound
		result.Msg = "The Tag name already exists"
		return
	}

	if tag, err := models.AddTag(tag.Name, tag.State, user.UserId); err != nil {
		result.Code = http.StatusInternalServerError
		result.Msg = err.Error()
	} else {
		result.Code = http.StatusOK
		result.Data = tag
	}
}

//修改文章标签
func EditTag(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)
	id := com.StrTo(c.Param("id")).MustInt64()
	var tag Tag
	_ = c.ShouldBindJSON(&tag)
	user := util.GetUser(c)

	valid := validation.Validation{}

	valid.Required(id, "id").Message("id不能为空")
	valid.MaxSize(tag.Name, 16, "name").Message("name最长为16字符")

	if valid.HasErrors() {
		result.Code = http.StatusNotFound
		result.Msg = valid.Errors[0].Message
		return
	}

	if !models.ExistTagByID(id) {
		result.Code = http.StatusNotFound
		result.Msg = "The Tag don't exist"
		return
	}

	data := make(map[string]interface{})
	data["modified_by"] = user.UserId
	if tag.Name != "" {
		data["name"] = tag.Name
	}
	if tag, err := models.EditTag(id, data); err != nil {
		result.Code = http.StatusInternalServerError
		result.Msg = err.Error()
	} else {
		result.Code = http.StatusOK
		result.Data = tag
	}
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)
	id := com.StrTo(c.Param("id")).MustInt64()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		result.Code = http.StatusNotFound
		result.Msg = valid.Errors[0].Message
		return
	}
	if models.ExistTagByID(id) {
		models.DeleteTag(id)
		result.Data = "success"
		result.Code = http.StatusOK
		return
	} else {
		result.Code = http.StatusInternalServerError
		result.Msg = "The Tag don't exist"
		return
	}
}
