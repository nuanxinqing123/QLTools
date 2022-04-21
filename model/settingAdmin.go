// -*- coding: utf-8 -*-
// @Time    : 2022/4/21 19:48
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : settingAdmin.go

package model

// WebSettings 网站配置模型
type WebSettings struct {
	Key   string `json:"key" gorm:"primaryKey" binding:"required"`
	Value string `json:"value"`
}
