// -*- coding: utf-8 -*-
// @Time    : 2022/4/24 19:33
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : conAdmin.go

package model

import (
	"gorm.io/gorm"
)

// OperationRecord 日志记录数据表
type OperationRecord struct {
	gorm.Model
	Operation string // 操作方式
	Journal   string // 记录日志
}

// TransferM 容器：迁移数据
type TransferM struct {
	IDOne int `json:"IDOne"`
	IDTwo int `json:"IDTwo"`
}

// CopyM 容器：复制数据
type CopyM struct {
	IDOne int `json:"IDOne"`
	IDTwo int `json:"IDTwo"`
}

// BackupM 容器：备份数据
type BackupM struct {
	IDOne int `json:"IDOne"`
}

// PanelAllEnv 面板全部变量数据
type PanelAllEnv struct {
	Code int      `json:"code"`
	Data []AllEnv `json:"data"`
}

type AllEnv struct {
	ID      int    `json:"ID"`
	Name    string `json:"name"`
	Value   string `json:"value"`
	Remarks string `json:"remarks"`
}
