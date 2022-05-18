// -*- coding: utf-8 -*-
// @Time    : 2022/4/7 16:39
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : panel.go

package controllers

import (
	"QLPanelTools/logic"
	"QLPanelTools/model"
	res "QLPanelTools/tools/response"
	val "QLPanelTools/tools/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// PanelAdd 新增变量名
func PanelAdd(c *gin.Context) {
	// 获取参数
	p := new(model.PanelData)
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
	resCode := logic.PanelAdd(p)
	switch resCode {
	case res.CodeStorageFailed:
		// 创建面板信息出错
		res.ResError(c, res.CodeStorageFailed)
	case res.CodeSuccess:
		res.ResSuccess(c, "面板信息创建成功")
	}
}

// PanelUp 修改变量名
func PanelUp(c *gin.Context) {
	// 获取参数
	p := new(model.UpPanelData)
	if err := c.ShouldBindJSON(&p); err != nil {
		fmt.Println(p)
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
	resCode := logic.PanelUpdate(p)
	switch resCode {
	case res.CodeSuccess:
		res.ResSuccess(c, "面板信息更新成功")
	}
}

// PanelDel 删除变量名
func PanelDel(c *gin.Context) {
	// 获取参数
	p := new(model.DelPanelData)
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
	resCode := logic.PanelDelete(p)
	switch resCode {
	case res.CodeSuccess:
		res.ResSuccess(c, "面板信息删除成功")
	}
}

// GetAllPanelData 获取面板全部信息
func GetAllPanelData(c *gin.Context) {
	// 处理业务
	env, resCode := logic.GetAllPanelData()
	switch resCode {
	case res.CodeSuccess:
		// 修改成功
		res.ResSuccess(c, env)
	}
}

// UpdatePanelEnvData 修改面板绑定变量
func UpdatePanelEnvData(c *gin.Context) {
	// 获取参数
	p := new(model.PanelEnvData)
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
	resCode := logic.UpdatePanelEnvData(p)
	switch resCode {
	case res.CodeSuccess:
		// 修改成功
		res.ResSuccess(c, "修改成功")
	}
}
