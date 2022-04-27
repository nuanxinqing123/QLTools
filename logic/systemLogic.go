// -*- coding: utf-8 -*-
// @Time    : 2022/4/23 15:33
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : systemLogic.go

package logic

import (
	_const "QLPanelTools/const"
	"QLPanelTools/model"
	"QLPanelTools/tools/requests"
	res "QLPanelTools/tools/response"
	"encoding/json"
	"go.uber.org/zap"
)

// CheckVersion 检查版本更新
func CheckVersion() (model.WebVer, res.ResCode) {
	// 版本号
	var v model.Ver
	var w model.WebVer
	// 获取仓库版本信息
	url := "https://version.6b7.xyz/qltools_version.json"
	r, err := requests.Requests("GET", url, "", "")
	if err != nil {
		return w, res.CodeServerBusy
	}
	// 序列化内容
	err = json.Unmarshal(r, &v)
	if err != nil {
		zap.L().Error(err.Error())
		return w, res.CodeServerBusy
	}

	if v.Version != _const.Version {
		w.Update = true
	} else {
		w.Update = false
	}
	w.Notice = v.Notice

	return w, res.CodeSuccess
}
