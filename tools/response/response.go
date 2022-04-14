// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 14:04
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : response.go

package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code": 1001, 		// 程序中的错误码
	"msg": "xxx", 		// 提示信息
	"data": {}			// 数据
}
*/

type Data struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// ResError 返回错误信息
func ResError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK,
		&Data{
			Code: code,
			Msg:  code.Msg(),
			Data: nil,
		})
}

// ResErrorWithMsg 自定义错误返回
func ResErrorWithMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK,
		&Data{
			Code: code,
			Msg:  msg,
			Data: nil,
		})
}

// ResSuccess 返回成功信息
func ResSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK,
		&Data{
			Code: CodeSuccess,
			Msg:  CodeSuccess.Msg(),
			Data: data,
		})
}
