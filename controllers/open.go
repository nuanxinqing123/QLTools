// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 14:17
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : open.go

package controllers

import (
	"QLPanelTools/logic"
	"QLPanelTools/model"
	"QLPanelTools/sqlite"
	res "QLPanelTools/tools/response"
	val "QLPanelTools/tools/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandle 注册请求
func SignUpHandle(c *gin.Context) {
	// 获取参数
	p := new(model.UserSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 参数校验
		zap.L().Error("SignUpHandle with invalid param", zap.Error(err))

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

	// 业务处理
	resCode := logic.SignUp(p)
	switch resCode {
	case res.CodeServerBusy:
		// 内部服务错误
		res.ResError(c, res.CodeServerBusy)
	case res.CodeSuccess:
		// 注册成功
		res.ResSuccess(c, "注册完成")
	case res.CodeRegistrationClosed:
		// 已关闭注册,禁用注册
		res.ResError(c, res.CodeRegistrationClosed)
	}
}

// SignInHandle 登录请求
func SignInHandle(c *gin.Context) {
	// 获取参数
	p := new(model.UserSignIn)
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
	token, resCode := logic.SignIn(p)
	switch resCode {
	case res.CodeEmailNotExist:
		go logic.AddIPAddr(c.ClientIP(), false)
		// 邮箱不存在
		res.ResError(c, res.CodeEmailNotExist)
	case res.CodeEmailFormatError:
		go logic.AddIPAddr(c.ClientIP(), false)
		// 邮箱格式错误
		res.ResError(c, res.CodeEmailFormatError)
	case res.CodeInvalidPassword:
		go logic.AddIPAddr(c.ClientIP(), false)
		// 邮箱或者密码错误
		res.ResError(c, res.CodeInvalidPassword)
	case res.CodeServerBusy:
		go logic.AddIPAddr(c.ClientIP(), false)
		// 生成Token出错
		res.ResError(c, res.CodeServerBusy)
	case res.CodeSuccess:
		go logic.AddIPAddr(c.ClientIP(), true)
		// 登录成功,返回Token
		res.ResSuccess(c, token)
	}
}

// CheckToken 检查Token是否有效
func CheckToken(c *gin.Context) {
	// 获取参数
	p := new(model.CheckToken)
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
	bol, resCode := logic.CheckToken(p)
	switch resCode {
	case res.CodeServerBusy:
		// 内部服务错误
		res.ResError(c, res.CodeServerBusy)
	case res.CodeInvalidToken:
		// Token已失效
		res.ResErrorWithMsg(c, res.CodeInvalidToken, "无效的Token")
	case res.CodeSuccess:
		// 上传成功
		res.ResSuccess(c, bol)
	}
}

// IndexData 可用服务
func IndexData(c *gin.Context) {
	// 处理业务
	resCode, data := logic.EnvData()
	switch resCode {
	case res.CodeDataError:
		res.ResErrorWithMsg(c, res.CodeDataError, "发生一点小意外，请刷新页面")
	case res.CodeCheckDataNotExist:
		// 获取数据不存在
		res.ResError(c, res.CodeCheckDataNotExist)
	case res.CodeServerBusy:
		// 内部服务错误
		res.ResError(c, res.CodeServerBusy)
	case res.CodeSuccess:
		// 登录成功,返回Token
		res.ResSuccess(c, data)
	}
}

// EnvADD 上传变量
func EnvADD(c *gin.Context) {
	// 获取参数
	p := new(model.EnvAdd)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 参数校验
		zap.L().Error("SignUpHandle with invalid param", zap.Error(err))

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
	// 查询请求IP是否受限
	resCode := logic.CheckIPIfItNormal(c.ClientIP())
	if resCode == res.CodeNumberDepletion {
		res.ResErrorWithMsg(c, res.CodeNumberDepletion, "今日提交已到达上限")
		return
	} else if resCode == res.CodeServerBusy {
		res.ResErrorWithMsg(c, res.CodeServerBusy, "服务繁忙,请稍后重试")
		return
	}
	// 业务处理
	resCode, msg := logic.EnvAdd(p)
	switch resCode {
	case res.CodeServerBusy:
		res.ResErrorWithMsg(c, res.CodeServerBusy, "服务繁忙,请稍后重试")
	case res.CodeStorageFailed:
		res.ResErrorWithMsg(c, res.CodeStorageFailed, msg)
	case res.CodeErrorOccurredInTheRequest:
		res.ResErrorWithMsg(c, res.CodeErrorOccurredInTheRequest, "提交服务器或变量名不在白名单")
	case res.CodeEnvDataMismatch:
		res.ResErrorWithMsg(c, res.CodeEnvDataMismatch, "上传内容不符合规定, 请检查后再提交")
	case res.CodeDataIsNull:
		res.ResErrorWithMsg(c, res.CodeDataIsNull, "上传内容能为空")
	case res.CodeLocationFull:
		res.ResErrorWithMsg(c, res.CodeLocationFull, "限额已满，禁止提交")
	case res.CodeNoDuplicateSubmission:
		res.ResErrorWithMsg(c, res.CodeNoDuplicateSubmission, "禁止提交重复数据")
	case res.CodeBlackListEnv:
		res.ResErrorWithMsg(c, res.CodeBlackListEnv, "变量已被管理员禁止提交")
	case res.CodeCustomError:
		// JS执行发生错误, 系统错误
		res.ResErrorWithMsg(c, res.CodeCustomError, "执行插件发生错误，错误原因："+msg)
	case res.CodeNoAdmittance:
		// 数据禁止通过
		res.ResErrorWithMsg(c, res.CodeNoAdmittance, msg)
	case res.CodeCDKError:
		res.ResErrorWithMsg(c, res.CodeCDKError, msg)
	case res.CodeSuccess:
		// 记录上传IP
		go sqlite.InsertSubmitRecord(c.ClientIP())
		// 更新CDK使用次数
		go sqlite.UpdateCDKAvailableTimes(p)
		// 上传成功
		res.ResSuccess(c, "上传成功")
	}
}

// CheckCDK CDK检查
func CheckCDK(c *gin.Context) {
	// 获取参数
	p := new(model.CheckCDK)
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
	resCode, str := logic.CheckCDK(p)
	switch resCode {
	case res.CodeSuccess:
		// 上传成功
		res.ResSuccess(c, str)
	}
}
