// -*- coding: utf-8 -*-
// @Time    : 2022/4/7 16:42
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : panelLogic.go

package logic

import (
	"QLPanelTools/server/model"
	"QLPanelTools/server/sqlite"
	"QLPanelTools/tools/panel"
	res "QLPanelTools/tools/response"
	"go.uber.org/zap"
)

// PanelAdd 面板信息存入数据库
func PanelAdd(p *model.PanelData) res.ResCode {
	// 保存进数据库
	err := sqlite.InsertPanelData(p)
	if err != nil {
		zap.L().Error("Error inserting database, err:", zap.Error(err))
		return res.CodeStorageFailed
	}

	return res.CodeSuccess
}

// PanelUpdate 更新面板信息
func PanelUpdate(p *model.UpPanelData) res.ResCode {
	// 更新数据库
	sqlite.UpdatePanelData(p)
	return res.CodeSuccess
}

// PanelDelete 删除面板信息
func PanelDelete(p *model.DelPanelData) res.ResCode {
	// 删除面板信息
	sqlite.DelPanelData(p)
	return res.CodeSuccess
}

// GetAllPanelData 获取面板全部信息
func GetAllPanelData() ([]model.QLPanel, res.ResCode) {
	// 获取信息
	panel := sqlite.GetPanelAllData()
	return panel, res.CodeSuccess
}

// UnbindPanelEnvData 解除所有面板变量绑定
func UnbindPanelEnvData() res.ResCode {
	// 获取所有面板信息
	panel := sqlite.GetPanelAllData()
	for i := 0; i < len(panel); i++ {
		panel[i].EnvBinding = ""
		// 保存数据
		sqlite.UnbindPanelEnvData(panel[i])
	}

	return res.CodeSuccess
}

// UpdatePanelEnvData 修改面板绑定变量
func UpdatePanelEnvData(p *model.PanelEnvData) res.ResCode {
	sqlite.UpdatePanelEnvData(p)
	return res.CodeSuccess
}

// UpdateAllPanelToken 解除所有面板变量绑定
func UpdateAllPanelToken() res.ResCode {
	// 获取所有面板信息
	panelData := sqlite.GetPanelAllData()
	for i := 0; i < len(panelData); i++ {
		// 更新Token
		go panel.GetPanelToken(panelData[i].URL, panelData[i].ClientID, panelData[i].ClientSecret)
	}

	return res.CodeSuccess
}
