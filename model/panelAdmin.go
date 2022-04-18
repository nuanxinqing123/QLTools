// -*- coding: utf-8 -*-
// @Time    : 2022/4/6 17:45
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : panelAdmin.go

package model

import "gorm.io/gorm"

// QLPanel QL面板数据
type QLPanel struct {
	gorm.Model
	PanelName    string `binding:"required"` // 面板名称
	URL          string `binding:"required"` // 面板连接地址
	ClientID     string `binding:"required"` // 面板Client_ID
	ClientSecret string `binding:"required"` // 面板Client_Secret
	Token        string // 面板Token
	Params       int    // 面板Params
	EnvBinding   string // 绑定变量
}

// PanelAll 全部面板信息
type PanelAll struct {
	ID           uint   `json:"ID"`
	PanelName    string `json:"name"`
	URL          string `json:"url"`
	ClientID     string `json:"id"`
	ClientSecret string `json:"secret"`
	EnvBinding   string `json:"envBinding"`
}

// PanelData 创建面板数据
type PanelData struct {
	Name   string `json:"name"`                      // 面板名称
	URL    string `json:"url" binding:"required"`    // 面板连接地址
	ID     string `json:"id" binding:"required"`     // 面板Client_ID
	Secret string `json:"secret" binding:"required"` // 面板Client_Secret
}

// UpPanelData 更新面板数据
type UpPanelData struct {
	UID    int    `json:"uid" binding:"required"`    // 数据库ID值
	Name   string `json:"name" binding:"required"`   // 面板名称
	URL    string `json:"url" binding:"required"`    // 面板连接地址
	ID     string `json:"id" binding:"required"`     // 面板Client_ID
	Secret string `json:"secret" binding:"required"` // 面板Client_Secret
}

// DelPanelData 删除面板数据
type DelPanelData struct {
	UID int `json:"uid" binding:"required"` // 数据库ID值
}

// PanelEnvData 修改面板绑定变量
type PanelEnvData struct {
	UID        int      `json:"uid" binding:"required"`        // 数据库ID值
	EnvBinding []string `json:"envBinding" binding:"required"` // 变量值
}

// EnvSData 可用服务器
type EnvSData struct {
	Count      int           `json:"count"`
	ServerData []envSData    `json:"serverData"`
	EnvData    []envNameData `json:"envData"`
}

type envSData struct {
	ID   int    `json:"ID"`
	Name string `json:"PanelName"`
}

// Token 面板Token数据
type Token struct {
	Code int `json:"code"`
	Data struct {
		Token      string `json:"token"`
		TokenType  string `json:"token_type"`
		Expiration int    `json:"expiration"`
	} `json:"data"`
}

type ResAdd struct {
	Code int `json:"code"`
}

// EnvData 面板变量数据
type EnvData struct {
	Code int       `json:"code"`
	Data []envData `json:"data"`
}

type envData struct {
	Name string `json:"name"`
}
