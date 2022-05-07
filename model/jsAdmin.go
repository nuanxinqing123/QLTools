// -*- coding: utf-8 -*-
// @Time    : 2022/5/6 16:16
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : jsAdmin.go

package model

// DeletePlugin 删除插件
type DeletePlugin struct {
	FileName string `json:"FileName"`
}

// FileInfo 读取插件信息
type FileInfo struct {
	FileName   string `json:"FileName"`
	FileIDName string `json:"FileIDName"`
}
