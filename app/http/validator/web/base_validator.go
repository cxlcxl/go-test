package validator_web

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"
	"goskeleton/app/global/variable"
)

var trans ut.Translator

// LoadValidatorLocal 初始化语言包
func LoadValidatorLocal() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(enT, zhT, enT)

		var o bool
		local := variable.ConfigYml.GetString("Validator.Local")
		trans, o = uni.GetTranslator(local)
		if !o {
			return errors.New("uni.GetTranslator failed")
		}
		// register translate
		// 注册翻译器
		var err error
		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = chTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = chTranslations.RegisterDefaultTranslations(v, trans)
		}

		if err != nil {
			variable.ZapLog.Error("初始化验证语言失败", zap.Error(err))
			return err
		}
		return nil
	}
	return nil
}

func Translate(err error) (errMsg string) {
	errs := err.(validator.ValidationErrors)
	for _, err := range errs {
		errMsg = err.Translate(trans)
		break
	}
	return
}
