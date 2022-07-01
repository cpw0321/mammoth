package internal

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func init() {
	//注册翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	//注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

// Translate 翻译错误信息
func Translate(err error) string {

	var result string

	errors := err.(validator.ValidationErrors)

	for _, err := range errors {
		result = result + " | " + err.Translate(trans)
	}
	return result
}
