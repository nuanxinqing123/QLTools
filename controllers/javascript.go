// -*- coding: utf-8 -*-
// @Time    : 2022/5/5 20:28
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : javascript.go

package controllers

import (
	"QLPanelTools/tools/goja"
	"fmt"
	"github.com/gin-gonic/gin"
)

// JavascriptTest 测试
func JavascriptTest(c *gin.Context) {
	js, i, s, err := goja.RunJS("test.js", "ceshi_env")
	if err != nil {
		return
	}
	fmt.Println(js)
	fmt.Println(i)
	fmt.Println(s)

	c.String(200, s)
}
