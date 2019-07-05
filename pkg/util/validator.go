package util

import (
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"sync"
)

type MyValidator struct {
	once     sync.Once
	validate *validator.Validate
}


func (v *MyValidator) ValidateStruct(obj interface{}) error {
	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	if valueType == reflect.Struct {
		v.lazyinit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}


func (v *MyValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *MyValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("valid")

	})
}
