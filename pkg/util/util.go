package util

import (
	"gin-blog/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

func GenerateErrors(obj interface{}, errors []*validation.Error, bindString string) (errs []string) {
	if len(errors) == 0 {
		return
	}
	valueObj := reflect.TypeOf(obj)
	for _, err := range errors{
		field, _ := valueObj.FieldByName(err.Field)
		key := field.Tag.Get(bindString)
		if key == "" {
			key = err.Field
		}
		errString := strings.Join([]string{key, err.Message}, " ")
		errs = append(errs, errString)
	}
	return
}

func ValidStructure(obj interface{}, c *gin.Context) bool {
	valid := validation.Validation{}
	ok, err := valid.Valid(obj)
	if err != nil {
		c.JSON(e.ERROR, gin.H{
			"msg": err,
		})
		return false
	}

	if !ok {
		errs := GenerateErrors(obj, valid.Errors, "json")
		c.JSON(e.InvalidParams, gin.H{
			"msg": strings.Join(errs, ", "),
		})
		c.Abort()
		return false
	}
	return true
}
