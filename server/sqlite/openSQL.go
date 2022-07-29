// -*- coding: utf-8 -*-
// @Time    : 2022/4/7 19:21
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : openSQL.go

package sqlite

import (
	"QLPanelTools/server/model"
)

// GetServerCount 获取可用服务器和ID值
func GetServerCount() []model.QLPanel {
	var s []model.QLPanel
	DB.Where("enable = ?", true).Find(&s)
	return s
}

// CheckServerDoesItExist 检查服务器是否存在
func CheckServerDoesItExist(id int) (bool, model.QLPanel) {
	var s model.QLPanel
	DB.First(&s, id)
	if s.ID == 0 {
		return false, s
	}

	return true, s
}

// CheckEnvNameDoesItExist 检查变量是否存在
func CheckEnvNameDoesItExist(name string) (bool, model.EnvName) {
	var e model.EnvName
	DB.Where("name = ?", name).First(&e)
	if e.ID == 0 {
		return false, e
	}

	return true, e
}
