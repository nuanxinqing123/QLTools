// -*- coding: utf-8 -*-
// @Time    : 2022/6/12 14:09
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : cdk.go

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

// GetDivisionCDKData CDK：以20条数据分割
func GetDivisionCDKData(c *gin.Context) {
	// 获取查询页码
	page := c.Query("page")
	resCode, data := logic.GetDivisionCDKData(page)

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// GetCDKData 获取CDK数据
func GetCDKData(c *gin.Context) {
	method := c.Query("method")
	var data []model.CDK
	var resCode res.ResCode
	if method == "all" {
		// 查询全部数据
		resCode, data = logic.GetCDKData(1)
	} else if method == "true" {
		// 查询启用数据
		resCode, data = logic.GetCDKData(2)
	} else if method == "false" {
		// 查询禁用数据
		resCode, data = logic.GetCDKData(3)
	} else {
		// 查询搜索数据
		s := c.Query("s")
		zap.L().Debug("【CDK搜索】值：" + s)
		if s == "" {
			res.ResErrorWithMsg(c, res.CodeInvalidParam, "请求数据不完整")
			return
		}
		resCode, data = logic.GetOneCDKData(s)
	}

	switch resCode {
	case res.CodeSuccess:
		// 查询成功
		res.ResSuccess(c, data)
	}
}

// CreateCDKData 批量生成CDK
func CreateCDKData(c *gin.Context) {
	// 获取参数
	p := new(model.CreateCDK)
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
	resCode := logic.CreateCDKData(p)
	switch resCode {
	case res.CodeServerBusy:
		res.ResErrorWithMsg(c, res.CodeServerBusy, "生成CDK失败，请检查日志获取报错信息")
	case res.CodeCreateFileError:
		res.ResErrorWithMsg(c, res.CodeCreateFileError, "创建CDK已写入数据库，但生成下载文件失败")
	case res.CodeSuccess:
		// 生成成功
		res.ResSuccess(c, "生成成功")
	}
}

// DownloadCDKData 下载CDK文件
func DownloadCDKData(c *gin.Context) {
	Filename := "CDK.txt"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", Filename))
	c.File("./" + Filename)
	go logic.DelCDKData()
}

// CDKState CDK全部启用/禁用
func CDKState(c *gin.Context) {
	// 获取参数
	p := new(model.UpdateStateCDK)
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
	resCode := logic.CDKState(p)
	switch resCode {
	case res.CodeSuccess:
		// 更新成功
		res.ResSuccess(c, "更新成功")
	}
}

// UpdateCDK 更新单条CDK数据
func UpdateCDK(c *gin.Context) {
	// 获取参数
	p := new(model.UpdateCDK)
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
	resCode := logic.UpdateCDK(p)
	switch resCode {
	case res.CodeSuccess:
		// 更新成功
		res.ResSuccess(c, "更新成功")
	}
}

// DelCDK 删除CDK数据
func DelCDK(c *gin.Context) {
	// 获取参数
	p := new(model.DelCDK)
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
	resCode := logic.DelCDK(p)
	switch resCode {
	case res.CodeSuccess:
		// 删除成功
		res.ResSuccess(c, "删除成功")
	}
}
