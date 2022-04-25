// -*- coding: utf-8 -*-
// @Time    : 2022/4/24 19:31
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : containerLogic.go

package logic

import (
	"QLPanelTools/model"
	"QLPanelTools/sqlite"
	"QLPanelTools/tools/panel"
	"QLPanelTools/tools/requests"
	res "QLPanelTools/tools/response"
	"encoding/json"
	"go.uber.org/zap"
	"strconv"
)

// Transfer 容器：迁移
func Transfer(p *model.TransferM) res.ResCode {
	// 根据ID查询服务器信息
	oneData := sqlite.GetPanelDataByID(p.IDOne)
	twoData := sqlite.GetPanelDataByID(p.IDTwo)
	// 检查白名单
	if oneData.URL == "" {
		return res.CodePanelNotWhitelisted
	}
	if twoData.URL == "" {
		return res.CodePanelNotWhitelisted
	}

	// 为了保证服务高可用性,强制更新面板Token
	panel.GetPanelToken(oneData.URL, oneData.ClientID, oneData.ClientSecret)
	panel.GetPanelToken(twoData.URL, twoData.ClientID, twoData.ClientSecret)

	// 重新获取Token
	oneData = sqlite.GetPanelDataByID(p.IDOne)
	twoData = sqlite.GetPanelDataByID(p.IDTwo)

	// 获取One面板全部信息
	zap.L().Debug("容器迁移：获取One面板全部信息")
	url := panel.StringHTTP(oneData.URL) + "/open/envs?searchValue=&t=" + strconv.Itoa(oneData.Params)
	allData, _ := requests.Requests("GET", url, "", oneData.Token)

	// 绑定数据
	var token model.PanelAllEnv
	err := json.Unmarshal(allData, &token)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	// 判断返回状态
	if token.Code != 200 {
		// 未授权或Token失效
		return res.CodeErrorOccurredInTheRequest
	}

	// 向B容器上传变量
	zap.L().Debug("容器迁移：向B容器上传变量")
	go EnvUp(token, twoData.URL, twoData.Token, twoData.Params, "迁移任务(上传)")

	// 获取A容器所有变量ID
	idGroup := `[`
	for i := 0; i < len(token.Data); i++ {
		idGroup = idGroup + strconv.Itoa(token.Data[i].ID) + `,`
	}
	idGroup = idGroup[:len(idGroup)-1]
	idGroup = idGroup + `]`

	// 删除A容器变量
	zap.L().Debug("容器迁移：删除A容器变量")
	EnvDel(idGroup, oneData.URL, oneData.Token, oneData.Params, "迁移任务(删除)")

	return res.CodeSuccess
}

// Copy 容器：复制
func Copy(p *model.CopyM) res.ResCode {
	// 根据ID查询服务器信息
	oneData := sqlite.GetPanelDataByID(p.IDOne)
	twoData := sqlite.GetPanelDataByID(p.IDTwo)
	// 检查白名单
	if oneData.URL == "" {
		return res.CodePanelNotWhitelisted
	}
	if twoData.URL == "" {
		return res.CodePanelNotWhitelisted
	}

	// 为了保证服务高可用性,强制更新面板Token
	panel.GetPanelToken(oneData.URL, oneData.ClientID, oneData.ClientSecret)
	panel.GetPanelToken(twoData.URL, twoData.ClientID, twoData.ClientSecret)

	// 重新获取Token
	oneData = sqlite.GetPanelDataByID(p.IDOne)
	twoData = sqlite.GetPanelDataByID(p.IDTwo)

	// 获取One面板全部信息
	zap.L().Debug("容器复制：获取One面板全部信息")
	url := panel.StringHTTP(oneData.URL) + "/open/envs?searchValue=&t=" + strconv.Itoa(oneData.Params)
	allData, _ := requests.Requests("GET", url, "", oneData.Token)

	// 绑定数据
	var token model.PanelAllEnv
	err := json.Unmarshal(allData, &token)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	// 判断返回状态
	if token.Code != 200 {
		// 未授权或Token失效
		return res.CodeErrorOccurredInTheRequest
	}

	// 向B容器上传变量
	zap.L().Debug("容器复制：向B容器上传变量")
	go EnvUp(token, twoData.URL, twoData.Token, twoData.Params, "复制任务(上传)")

	return res.CodeSuccess
}

// Backup 容器：备份
func Backup(p *model.BackupM) {

}

// Restore 容器：恢复
func Restore(p *model.RestoreM) {

}

// EnvUp 变量上传
func EnvUp(p model.PanelAllEnv, url, token string, params int, journal string) {
	// 获取数量
	for i := 0; i < len(p.Data); i++ {
		var data string
		URL := panel.StringHTTP(url) + "/open/envs?t=" + strconv.Itoa(params)
		// 上传
		data = `[{"value": "` + p.Data[i].Value + `","name": "` + p.Data[i].Name + `","remarks": "` + p.Data[i].Remarks + `"}]`
		// 执行上传任务
		_, err := requests.Requests("POST", URL, data, token)
		if err != nil {
			// 记录错误
			sqlite.RecordingError(journal, err.Error())
		}
	}
}

// EnvDel 变量删除
func EnvDel(p string, url, token string, params int, journal string) {
	URL := panel.StringHTTP(url) + "/open/envs?t=" + strconv.Itoa(params)
	// 执行删除任务
	_, err := requests.Requests("DELETE", URL, p, token)
	if err != nil {
		// 记录错误
		sqlite.RecordingError(journal, err.Error())
	}
}
