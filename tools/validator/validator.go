// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 13:14
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : validator.go

package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

// Trans 定义全局翻译器
var Trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改框架中的Validator引擎属性
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取JSON Tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 HTTP 请求头的 “Accept-Language”
		var ok bool
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			err = enTrans.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTrans.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTrans.RegisterDefaultTranslations(v, Trans)
		}
		return
	}
	return
}

// RemoveTopStruct 分割结构体名称
func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
