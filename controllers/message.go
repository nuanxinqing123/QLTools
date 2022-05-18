// -*- coding: utf-8 -*-
// @Time    : 2022/5/17 12:25
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : message.go

package controllers

import (
	"QLPanelTools/logic"
	"QLPanelTools/model"
	res "QLPanelTools/tools/response"
	val "QLPanelTools/tools/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// GetEmailData 获取邮件信息
func GetEmailData(c *gin.Context) {
	resCode, data := logic.GetEmailData()
	switch resCode {
	case res.CodeSuccess:
		res.ResSuccess(c, data)
	}
}

// SendTestEmail 测试发送邮件
func SendTestEmail(c *gin.Context) {
	// 获取参数
	p := new(model.TestEmail)

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

	resCode := logic.TestEmailSend(p)
	switch resCode {
	case res.CodeServerBusy:
		res.ResErrorWithMsg(c, res.CodeServerBusy, "测试发送邮件失败，具体原因请查看日志记录")
	case res.CodeSuccess:
		res.ResSuccess(c, "测试邮件发送成功")
	}
}

// UpdateEmailSet 修改邮件配置
func UpdateEmailSet(c *gin.Context) {
	// 获取参数
	p := new(model.UpdateEmail)
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

	resCode := logic.UpdateEmailSet(p)
	switch resCode {
	case res.CodeSuccess:
		res.ResSuccess(c, "邮件服务信息修改成功")
	}
}
