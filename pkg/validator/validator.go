package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var (
	Trans ut.Translator
	vt    *validator.Validate
)

func InitValidator() {
	var ok bool
	vt, ok = binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	zhHans := zh_Hans.New()
	uni := ut.New(zhHans)

	Trans, _ = uni.GetTranslator("")
	_ = zh.RegisterDefaultTranslations(vt, Trans)

	vt.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("label")
	})

	registerCustom("identity", "{0}必须是1或2", func(fl validator.FieldLevel) bool {
		val := fl.Field().Int()
		return val == 1 || val == 2
	})
}

func registerCustom(tag string, transText string, fn validator.Func) {
	_ = vt.RegisterValidation(tag, fn)

	_ = vt.RegisterTranslation(tag, Trans, func(ut ut.Translator) error {
		return ut.Add(tag, transText, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T(fe.Tag(), fe.Field())
		if err != nil {
			return fe.Error()
		}
		return t
	})
}
