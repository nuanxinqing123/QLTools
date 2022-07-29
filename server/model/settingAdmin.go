// -*- coding: utf-8 -*-
// @Time    : 2022/4/21 19:48
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : settingAdmin.go

package model

import "gorm.io/gorm"

// WebSettings 网站配置模型
type WebSettings struct {
	Key   string `json:"key" gorm:"primaryKey" binding:"required"`
	Value string `json:"value"`
}

// IPSubmitRecord IP提交记录模型
type IPSubmitRecord struct {
	gorm.Model
	SubmitTime string `json:"submit_time" binding:"required"` // 提交时间（格式：2022-04-17）
	IPAddress  string `json:"ip_address" binding:"required"`  // 提交IP
}
