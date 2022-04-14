// -*- coding: utf-8 -*-
// @Time    : 2022/4/6 17:04
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : envLogic.go

package logic

import (
	"QLPanelTools/model"
	"QLPanelTools/sqlite"
	res "QLPanelTools/tools/response"
)

// EnvNameAdd 创建变量名
func EnvNameAdd(p *model.EnvNameAdd) res.ResCode {
	// 检查变量名是否已存在
	result := sqlite.CheckEnvName(p.EnvName)
	if result == true {
		return res.CodeEnvNameExist
	}

	// 不存在, 创建新的变量名
	err := sqlite.AddEnvName(p)
	if err != nil {
		return res.CodeStorageFailed
	}
	return res.CodeSuccess
}

// EnvNameUpdate 更新变量名
func EnvNameUpdate(p *model.EnvNameUp) res.ResCode {
	// 更新数据库
	sqlite.UpdateEnvName(p)
	return res.CodeSuccess
}

// EnvNameDel 删除变量名
func EnvNameDel(p *model.EnvNameDel) res.ResCode {
	// 删除变量名
	sqlite.DelEnvName(p)
	return res.CodeSuccess
}

// GetAllEnvData 获取变量全部信息
func GetAllEnvData() ([]*model.EnvName, res.ResCode) {
	// 获取信息
	env := sqlite.GetEnvNameAll()
	return env, res.CodeSuccess
}
