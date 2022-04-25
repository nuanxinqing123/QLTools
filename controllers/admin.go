// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 15:06
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : admin.go

package controllers

import (
	"QLPanelTools/logic"
	"QLPanelTools/model"
	"QLPanelTools/tools/panel"
	res "QLPanelTools/tools/response"
	val "QLPanelTools/tools/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

func AdminTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Code": 200,
	})
}

// GetPanelToken 面板连接测试
func GetPanelToken(c *gin.Context) {
	// 获取参数
	p := new(model.PanelData)
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
	resCode, token := panel.TestGetPanelToken(p.URL, p.ID, p.Secret)
	if resCode == res.CodeServerBusy {
		// 内部服务错误
		res.ResError(c, res.CodeServerBusy)
		return
	}

	if token.Code != 200 {
		// 授权错误
		res.ResErrorWithMsg(c, res.CodeDataError, "client_id或client_secret有误")
		return
	}

	res.ResSuccess(c, "面板连接测试成功")
	return
}

// ReAdminPwd 修改管理员密码
func ReAdminPwd(c *gin.Context) {
	// 获取参数
	p := new(model.ReAdminPwd)
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
	bol, resCode := logic.RePwd(p)
	switch resCode {
	case res.CodeOldPassWordError:
		// 旧密码错误
		res.ResError(c, res.CodeOldPassWordError)
	case res.CodeServerBusy:
		// 保存修改错误
		res.ResError(c, res.CodeServerBusy)
	case res.CodeSuccess:
		// 修改成功
		res.ResSuccess(c, bol)
	}
}

// GetIPInfo 获取近十次登录
func GetIPInfo(c *gin.Context) {
	// 处理业务
	bol, resCode := logic.GetIPInfo()
	switch resCode {
	case res.CodeSuccess:
		// 修改成功
		res.ResSuccess(c, bol)
	}
}

// GetAdminInfo 获取管理员信息
func GetAdminInfo(c *gin.Context) {
	info, resCode := logic.GetAdminInfo()
	switch resCode {
	case res.CodeSuccess:
		// 修改成功
		res.ResSuccess(c, info)
	}
}
