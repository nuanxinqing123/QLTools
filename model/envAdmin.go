// -*- coding: utf-8 -*-
// @Time    : 2022/4/6 16:45
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : envAdmin.go

package model

import "gorm.io/gorm"

// EnvName 变量名
type EnvName struct {
	gorm.Model
	Name     string `binding:"required"` // 环境变量名称
	Quantity int    `binding:"required"` // 环境变量数量上限
	Regex    string // 环境变量匹配正则
	Mode     int    `binding:"required"` // 环境变量模式[1：新建模式、2：合并模式]
	Division string // 环境变量分隔符（合并模式）
}

// EnvNameAdd 新增变量名
type EnvNameAdd struct {
	EnvName     string `json:"envName" binding:"required"`
	EnvQuantity int    `json:"envQuantity" binding:"required"`
	EnvRegex    string `json:"envRegex"`
	EnvMode     int    `json:"envMode" binding:"required"`
	EnvDivision string `json:"envDivision"`
}

// EnvNameUp 修改变量名
type EnvNameUp struct {
	EnvID       int    `json:"envID" binding:"required"`
	EnvName     string `json:"envName" binding:"required"`
	EnvQuantity int    `json:"envQuantity" binding:"required"`
	EnvRegex    string `json:"envRegex"`
	EnvMode     int    `json:"envMode" binding:"required"`
	EnvDivision string `json:"envDivision"`
}

// EnvNameDel 删除变量名
type EnvNameDel struct {
	EnvID int `json:"envID" binding:"required"`
}

// envNameData 变量数据
type envNameData struct {
	// 变量名称
	Name string `json:"name"`
	// 变量剩余限额
	Quantity int `json:"quantity"`
}

// EnvAdd 上传变量
type EnvAdd struct {
	// 服务器ID
	ServerID int `json:"serverID" binding:"required"`
	// 变量名
	EnvName string `json:"envName"  binding:"required"`
	// 变量值
	EnvData string `json:"envData"  binding:"required"`
	// 备注
	EnvRemarks string `json:"envRemarks"`
}
