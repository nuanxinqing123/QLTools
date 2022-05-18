// -*- coding: utf-8 -*-
// @Time    : 2022/4/23 15:30
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : system.go

package controllers

import (
	"QLPanelTools/logic"
	res "QLPanelTools/tools/response"
	"github.com/gin-gonic/gin"
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
	resCode, _ := logic.UpdateSoftware()
	switch resCode {
	case res.CodeServerBusy:
		res.ResError(c, res.CodeServerBusy)
	case res.CodeSuccess:
		// 获取成功
		res.ResSuccess(c, res.CodeSuccess)
	}
}
