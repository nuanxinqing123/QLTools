// -*- coding: utf-8 -*-
// @Time    : 2022/4/23 15:30
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : system.go

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

// CheckVersion 检查版本更新
func CheckVersion(c *gin.Context) {
	v, resCode := logic.CheckVersion()
	switch resCode {
	case res.CodeServerBusy:
		res.ResError(c, res.CodeServerBusy)
	case res.CodeSuccess:
		// 获取成功
		res.ResSuccess(c, v)
	}
}

// UpdateSoftware 更新软件
func UpdateSoftware(c *gin.Context) {
	// 获取参数
	p := new(model.SoftWareGOOS)
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

	resCode, txt := logic.UpdateSoftware(p)
	switch resCode {
	case res.CodeUpdateServerBusy:
		res.ResErrorWithMsg(c, res.CodeUpdateServerBusy, txt)
	case res.CodeSuccess:
		// 获取成功
		res.ResSuccess(c, txt)
	}
}
