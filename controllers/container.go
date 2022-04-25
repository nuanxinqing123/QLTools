// -*- coding: utf-8 -*-
// @Time    : 2022/4/24 19:27
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : container.go

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

// Transfer 容器：迁移
func Transfer(c *gin.Context) {
	// 获取参数
	p := new(model.TransferM)
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
	logic.Transfer(p)
	res.ResSuccess(c, "操作已进入任务队列, 请稍后前往青龙面板查看结果")
}

// Copy 容器：复制
func Copy(c *gin.Context) {
	// 获取参数
	p := new(model.CopyM)
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
	logic.Copy(p)
	res.ResSuccess(c, "操作已进入任务队列, 请稍后前往青龙面板查看结果")
}

// Backup 容器：备份
func Backup(c *gin.Context) {

	res.ResSuccess(c, "操作已进入任务队列, 请稍后前往青龙面板查看结果")
}

// Restore 容器：恢复
func Restore(c *gin.Context) {

	res.ResSuccess(c, "操作已进入任务队列, 请稍后前往青龙面板查看结果")
}
