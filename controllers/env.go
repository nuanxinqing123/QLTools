// -*- coding: utf-8 -*-
// @Time    : 2022/4/6 16:48
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : env.go

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

// EnvNameAdd 新增变量名
func EnvNameAdd(c *gin.Context) {
	// 获取参数
	p := new(model.EnvNameAdd)
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
	resCode := logic.EnvNameAdd(p)
	switch resCode {
	case res.CodeStorageFailed:
		// 变量名创建失败
		res.ResError(c, res.CodeStorageFailed)
	case res.CodeEnvNameExist:
		// 变量名已存在
		res.ResError(c, res.CodeEnvNameExist)
	case res.CodeSuccess:
		res.ResSuccess(c, "变量名创建成功")
	}
}

// EnvNameUp 修改变量名
func EnvNameUp(c *gin.Context) {
	// 获取参数
	p := new(model.EnvNameUp)
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
	resCode := logic.EnvNameUpdate(p)
	switch resCode {
	case res.CodeSuccess:
		res.ResSuccess(c, "变量名更新成功")
	}
}

// EnvNameDel 删除变量名
func EnvNameDel(c *gin.Context) {
	// 获取参数
	p := new(model.EnvNameDel)
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
	resCode := logic.EnvNameDel(p)
	switch resCode {
	case res.CodeSuccess:
		res.ResSuccess(c, "变量删除新成功")
	}
}

// GetAllEnvData 获取变量全部信息
func GetAllEnvData(c *gin.Context) {
	// 处理业务
	env, resCode := logic.GetAllEnvData()
	switch resCode {
	case res.CodeSuccess:
		// 修改成功
		res.ResSuccess(c, env)
	}
}

// GetPanelToken 获取面板Token
//func GetPanelToken(c *gin.Context) {
//	// 获取参数
//	p := new(model.PanelData)
//	if err := c.ShouldBindJSON(&p); err != nil {
//		// 参数校验
//		zap.L().Error("SignInHandle with invalid param", zap.Error(err))
//
//		// 判断err是不是validator.ValidationErrors类型
//		errs, ok := err.(validator.ValidationErrors)
//		if !ok {
//			res.ResError(c, res.CodeInvalidParam)
//			return
//		}
//
//		// 翻译错误
//		res.ResErrorWithMsg(c, res.CodeInvalidParam, val.RemoveTopStruct(errs.Translate(val.Trans)))
//		return
//	}
//
//	// 处理业务
//	resCode, token := requests.GetPanelToken(p.URL, p.ID, p.Secret)
//	switch resCode {
//	case res.CodeSuccess:
//		// 注册成功
//		res.ResSuccess(c, token)
//	case res.CodeServerBusy:
//		// 服务出现异常
//		res.ResError(c, res.CodeServerBusy)
//	}
//}
