// -*- coding: utf-8 -*-
// @Time    : 2022/4/21 19:47
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : setting.go

package controllers

import (
	"QLPanelTools/server/logic"
	"QLPanelTools/server/model"
	res "QLPanelTools/tools/response"
	val "QLPanelTools/tools/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// GetSetting 获取单个配置
func GetSetting(c *gin.Context) {
	key := c.Query("key")
	data, resCode := logic.GetSetting(key)
	switch resCode {
	case res.CodeServerBusy:
		// 越权
		res.ResErrorWithMsg(c, res.CodeServerBusy, "获取内容为空")
	case res.CodeSuccess:
		// 获取成功
		res.ResSuccess(c, data)
	}
}

// GetSettings 获取全部配置
func GetSettings(c *gin.Context) {
	data, resCode := logic.GetSettings()
	switch resCode {
	case res.CodeServerBusy:
		res.ResError(c, res.CodeServerBusy)
	case res.CodeSuccess:
		// 获取成功
		res.ResSuccess(c, data)
	}
}

// SaveSettings 保存网站信息
func SaveSettings(c *gin.Context) {
	p := new([]model.WebSettings)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 参数校验
		zap.L().Error("SignInHandle with invalid param", zap.Error(err))

		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.ResError(c, res.CodeInvalidParam)
			return
		}

		// 翻译错误
		res.ResErrorWithMsg(c, res.CodeInvalidParam, val.RemoveTopStruct(errs.Translate(val.Trans)))
		return
	}

	// 处理业务
	resCode := logic.SaveSettings(p)
	switch resCode {
	case res.CodeServerBusy:
		res.ResError(c, res.CodeServerBusy)
	case res.CodeSuccess:
		// 修改成功
		res.ResSuccess(c, "保存成功")
	}
}
