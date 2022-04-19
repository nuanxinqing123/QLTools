// -*- coding: utf-8 -*-
// @Time    : 2022/4/7 19:19
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : openLogic.go

package logic

import (
	"QLPanelTools/model"
	"QLPanelTools/sqlite"
	"QLPanelTools/tools/panel"
	"QLPanelTools/tools/requests"
	res "QLPanelTools/tools/response"
	"encoding/json"
	"go.uber.org/zap"
	"regexp"
	"strconv"
)

// EnvData 获取all data
func EnvData() (res.ResCode, model.EnvSData) {
	var sd model.EnvSData
	// 获取可用服务器
	sData, count := sqlite.GetServerCount()
	data, err := json.Marshal(sData)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy, sd
	}
	// 转化JSON脱敏
	err = json.Unmarshal(data, &sd.ServerData)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy, sd
	}
	sd.Count = count

	// 获取变量数据
	envData := sqlite.GetEnvNameAll()
	eData, err := json.Marshal(envData)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy, sd
	}
	err = json.Unmarshal(eData, &sd.EnvData)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy, sd
	}

	// 获取面板信息
	resCode, panelData := sqlite.GetPanelData()
	if resCode == res.CodeCheckDataNotExist {
		return res.CodeCheckDataNotExist, sd
	}

	// 获取面板已存在变量数量
	url := panel.StringHTTP(panelData.URL) + "/open/envs?searchValue=&t=" + strconv.Itoa(panelData.Params)
	allData, err := requests.Requests("GET", url, "", panelData.Token)
	if err != nil {
		return res.CodeServerBusy, sd
	}
	var token model.EnvData
	err = json.Unmarshal(allData, &token)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy, sd
	}

	if token.Code != 200 {
		// 尝试获取授权
		go panel.GetPanelToken(panelData.URL, panelData.ClientID, panelData.ClientSecret)

		// 未授权或Token失效
		return res.CodeDataError, sd
	}

	// 如果面板变量为NULL
	if len(token.Data) == 0 {
		return res.CodeEnvIsNull, sd
	}

	// 计算变量剩余限额
	for x := 0; x < len(sd.EnvData); x++ {
		for i := 0; i < len(token.Data); i++ {
			if token.Data[i].Name == sd.EnvData[x].Name {
				sd.EnvData[x].Quantity = sd.EnvData[x].Quantity - 1
			}
		}
	}

	return res.CodeSuccess, sd
}

// EnvAdd 添加变量
func EnvAdd(p *model.EnvAdd) res.ResCode {
	var token model.ResAdd
	// 校验服务器ID
	result, sData := sqlite.CheckServerDoesItExist(p.ServerID)
	if result != true {
		// 服务器不存在
		return res.CodeErrorOccurredInTheRequest
	}

	// 校验变量名是否存在
	result, eData := sqlite.CheckEnvNameDoesItExist(p.EnvName)
	if result != true {
		// 变量不存在
		return res.CodeErrorOccurredInTheRequest
	}

	if p.EnvData == "" {
		return res.CodeDataIsNull
	}

	// 正则处理
	if eData.Regex != "" {
		// 需要处理正则
		reg := regexp.MustCompile(eData.Regex)
		// 匹配内容
		if reg != nil {
			s := reg.FindAllStringSubmatch(p.EnvData, -1)
			if len(s) == 0 {
				return res.CodeEnvDataMismatch
			}
		} else {
			return res.CodeServerBusy
		}
	}

	// 校验变量配额
	c := CalculateQuantity(p.ServerID, p.EnvName)
	if c > 1999 {
		return res.CodeServerBusy
	} else if c <= 0 {
		return res.CodeLocationFull
	}

	// 提交到服务器
	url := panel.StringHTTP(sData.URL) + "/open/envs?t=" + strconv.Itoa(sData.Params)
	data := `[{"value": "` + p.EnvData + `","name": "` + p.EnvName + `","remarks": "` + p.EnvRemarks + `"}]`

	r, err := requests.Requests("POST", url, data, sData.Token)
	if err != nil {
		return res.CodeServerBusy
	}

	// 序列化内容
	err = json.Unmarshal(r, &token)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	if token.Code != 200 {
		// 尝试更新Token
		go panel.GetPanelToken(sData.URL, sData.ClientID, sData.ClientSecret)
		return res.CodeStorageFailed
	}

	return res.CodeSuccess
}

// CalculateQuantity 计算变量剩余位置
func CalculateQuantity(id int, name string) int {
	// 获取变量数据
	count := sqlite.GetEnvNameCount(name)

	// 获取容器信息
	sData := sqlite.GetPanelDataByID(id)

	// 获取面板已存在变量数量
	url := panel.StringHTTP(sData.URL) + "/open/envs?searchValue=&t=" + strconv.Itoa(sData.Params)
	allData, err := requests.Requests("GET", url, "", sData.Token)
	if err != nil {
		return res.CodeServerBusy
	}
	var token model.EnvData
	err = json.Unmarshal(allData, &token)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	// 计算变量剩余限额
	c := count
	for i := 0; i < len(token.Data); i++ {
		if token.Data[i].Name == name {
			c--
		}
	}

	return c
}
