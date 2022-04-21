// -*- coding: utf-8 -*-
// @Time    : 2022/4/21 19:56
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : settingLogic.go

package logic

import (
	"QLPanelTools/model"
	"QLPanelTools/sqlite"
	res "QLPanelTools/tools/response"
	"go.uber.org/zap"
)

// GetSetting 获取一个配置信息
func GetSetting(name string) (model.WebSettings, res.ResCode) {
	data, err := sqlite.GetSetting(name)
	if err != nil {
		zap.L().Error(err.Error())
		return data, res.CodeServerBusy
	}

	return data, res.CodeSuccess
}

// GetSettings 获取所有配置信息
func GetSettings() (interface{}, res.ResCode) {
	data, err := sqlite.GetSettings()
	if err != nil {
		zap.L().Error(err.Error())
		return nil, res.CodeServerBusy
	}

	return data, res.CodeSuccess
}

// SaveSettings 保存网站信息
func SaveSettings(p *[]model.WebSettings) res.ResCode {
	if err := sqlite.SaveSettings(p); err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	return res.CodeSuccess
}
