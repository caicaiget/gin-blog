package v1

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state = 1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS
	pageNum, pageSize := util.GetPage(c)
	data["lists"] = models.GetTags(pageNum, pageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}


// @Summary Get multiple article tags
// @Produce json
// @Param name query string false "name"
// @Param state query int false "state"
// @Param state query int false "createdBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":500,"data":{},"msg":"error"}"
// @Router /api/v1/tags [get]
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "1")).MustInt()
	user, _ := c.Get("user")
	userId := com.StrTo(reflect.ValueOf(user).FieldByName("ID").String()).MustInt()

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	if valid.HasErrors() {
		c.JSON(e.InvalidParams, gin.H{
			"msg": valid.Errors[0].Message,
		})
		return
	}

	if models.ExistTagByName(name) {
		c.JSON(e.ErrorNotExistTag, gin.H{
			"msg": e.GetMsg(e.ErrorNotExistTag),
		})
		return
	}

	if tag, err := models.AddTag(name, state, userId); err != nil {
		c.JSON(e.ERROR, gin.H{
			"msg": err,
		})
	} else {
		c.JSON(e.SUCCESS, gin.H{
			"data": tag,
		})
	}
}

//修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modifiedBy")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modifiedBy").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modifiedBy").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.InvalidParams
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ErrorExistTag
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.InvalidParams
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ErrorNotExistTag
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
