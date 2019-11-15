package util

import (
	"errors"
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

func ValidStructure(obj interface{}, c *gin.Context) (err error) {
	valid := validation.Validation{}
	ok, err := valid.Valid(obj)
	if err != nil {
		return err
	}

	if !ok {
		errs := GenerateErrors(obj, valid.Errors, "json")
		return errors.New(strings.Join(errs, ", "))
	}
	return nil
}
