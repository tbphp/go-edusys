package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var Trans ut.Translator

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhHans := zh_Hans.New()
		uni := ut.New(zhHans)

		Trans, _ = uni.GetTranslator("")
		_ = zh.RegisterDefaultTranslations(v, Trans)

		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("label")
		})
	}
}
