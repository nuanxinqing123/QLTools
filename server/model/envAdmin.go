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
	Name        string // 环境变量名称
	NameRemarks string // 环境变量名称备注
	Quantity    int    // 环境变量数量上限
	Regex       string // 环境变量匹配正则
	Mode        int    // 环境变量模式[1：新建模式、2：合并模式、3、更新模式]
	Division    string // 环境变量分隔符（合并模式）
	ReUpdate    string // 环境变量更新匹配正则（更新模式）
	IsPlugin    bool   // 环境变量是否使用插件
	PluginName  string // 绑定的插件名称
	IsCDK       bool   // 环境变量是否绑定CDK
}

// EnvNameAdd 新增变量名
type EnvNameAdd struct {
	EnvName        string `json:"envName" binding:"required"`
	EnvNameRemarks string `json:"envNameRemarks"`
	EnvQuantity    int    `json:"envQuantity" binding:"required"`
	EnvRegex       string `json:"envRegex"`
	EnvMode        int    `json:"envMode" binding:"required"`
	EnvDivision    string `json:"envDivision"`
	EnvReUpdate    string `json:"envReUpdate"`
	EnvIsPlugin    bool   `json:"envIsPlugin"`
	EnvPluginName  string `json:"envPluginName"`
	EnvIsCDK       bool   `json:"envIsCDK"`
}

// EnvNameUp 修改变量名
type EnvNameUp struct {
	EnvID          int    `json:"envID" binding:"required"`
	EnvName        string `json:"envName" binding:"required"`
	EnvNameRemarks string `json:"envNameRemarks"`
	EnvQuantity    int    `json:"envQuantity" binding:"required"`
	EnvRegex       string `json:"envRegex"`
	EnvMode        int    `json:"envMode" binding:"required"`
	EnvDivision    string `json:"envDivision"`
	EnvReUpdate    string `json:"envReUpdate"`
	EnvIsPlugin    bool   `json:"envIsPlugin"`
	EnvPluginName  string `json:"envPluginName"`
	EnvIsCDK       bool   `json:"envIsCDK"`
}

// EnvNameDel 删除变量名
type EnvNameDel struct {
	EnvID int `json:"envID" binding:"required"`
}

// envNameData 变量数据
type envNameData struct {
	// 变量名称
	Name string `json:"name"`
	// 变量备注
	NameRemarks string `json:"nameRemarks"`
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
	// CDK
	EnvCDK string `json:"envCDK"`
}
