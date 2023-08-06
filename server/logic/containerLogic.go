// -*- coding: utf-8 -*-
// @Time    : 2022/4/24 19:31
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : containerLogic.go

package logic

import (
	"QLPanelTools/server/model"
	"QLPanelTools/server/sqlite"
	"QLPanelTools/tools/panel"
	"QLPanelTools/tools/requests"
	res "QLPanelTools/tools/response"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
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
	time.Sleep(time.Second)

	// 重新获取Token
	one := sqlite.GetPanelDataByID(p.IDOne)
	two := sqlite.GetPanelDataByID(p.IDTwo)

	// 获取One面板全部信息
	zap.L().Debug("容器迁移：获取One面板全部信息")
	url := panel.StringHTTP(one.URL) + "/open/envs?searchValue=&t=" + strconv.Itoa(one.Params)
	allData, _ := requests.Requests("GET", url, "", one.Token)

	// 绑定数据
	var token model.PanelAllEnv
	err := json.Unmarshal(allData, &token)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	// 向B容器上传变量
	zap.L().Debug("容器迁移：向B容器上传变量")
	go EnvUp(token, two.URL, two.Token, two.Params, "迁移任务(上传)")

	// 获取A容器所有变量ID
	idGroup := `[`
	for i := 0; i < len(token.Data); i++ {
		idGroup = idGroup + strconv.Itoa(token.Data[i].ID) + `,`
	}
	idGroup = idGroup[:len(idGroup)-1]
	idGroup = idGroup + `]`

	// 删除A容器变量
	zap.L().Debug("容器迁移：删除A容器变量")
	EnvDel(idGroup, one.URL, one.Token, one.Params, "迁移任务(删除)")

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
	time.Sleep(time.Second)

	// 重新获取Token
	one := sqlite.GetPanelDataByID(p.IDOne)
	two := sqlite.GetPanelDataByID(p.IDTwo)

	// 获取One面板全部信息
	zap.L().Debug("容器复制：获取One面板全部信息")
	url := panel.StringHTTP(one.URL) + "/open/envs?searchValue=&t=" + strconv.Itoa(one.Params)
	allData, _ := requests.Requests("GET", url, "", one.Token)

	// 绑定数据
	var token model.PanelAllEnv
	err := json.Unmarshal(allData, &token)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	// 向B容器上传变量
	zap.L().Debug("容器复制：向B容器上传变量")
	go EnvUp(token, two.URL, two.Token, two.Params, "复制任务(上传)")

	return res.CodeSuccess
}

// Backup 容器：备份
func Backup(p *model.BackupM) res.ResCode {
	// 根据ID查询服务器信息
	one := sqlite.GetPanelDataByID(p.IDOne)

	// 检查白名单
	if one.URL == "" {
		return res.CodePanelNotWhitelisted
	}

	// 为了保证服务高可用性,强制更新面板Token
	panel.GetPanelToken(one.URL, one.ClientID, one.ClientSecret)
	time.Sleep(time.Second)

	// 重新获取Token
	server := sqlite.GetPanelDataByID(p.IDOne)

	// 获取One面板全部信息
	zap.L().Debug("容器备份：获取面板全部信息")
	url := panel.StringHTTP(server.URL) + "/open/envs?searchValue=&t=" + strconv.Itoa(server.Params)
	allData, _ := requests.Requests("GET", url, "", server.Token)

	// 绑定数据
	var token model.PanelAllEnv
	err := json.Unmarshal(allData, &token)
	if err != nil {
		// 记录错误
		sqlite.RecordingError("备份任务", err.Error())
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	// 创建JSON文件
	_, err = os.Create("backup.json")
	if err != nil {
		// 记录错误
		sqlite.RecordingError("备份任务", err.Error())
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	// 打开JSON文件
	f, err := os.Open("backup.json")
	if err != nil {
		// 记录错误
		sqlite.RecordingError("备份任务", err.Error())
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}
	defer func(f *os.File) {
		err = f.Close()
		if err != nil {
			// 记录错误
			sqlite.RecordingError("备份任务", err.Error())
			zap.L().Error(err.Error())
		}
	}(f)

	// 序列化数据
	b, err := json.Marshal(token.Data)
	if err != nil {
		// 记录错误
		sqlite.RecordingError("备份任务", err.Error())
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	// 保存数据
	err = ioutil.WriteFile("backup.json", b, 0777)
	if err != nil {
		// 记录错误
		sqlite.RecordingError("备份任务", err.Error())
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	return res.CodeSuccess
}

// Restore 容器：恢复
func Restore(sID string) res.ResCode {
	// 根据ID查询服务器信息
	iID, err := strconv.Atoi(sID)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}
	one := sqlite.GetPanelDataByID(iID)

	// 检查白名单
	if one.URL == "" {
		return res.CodePanelNotWhitelisted
	}

	// 为了保证服务高可用性,强制更新面板Token
	panel.GetPanelToken(one.URL, one.ClientID, one.ClientSecret)
	time.Sleep(time.Second)

	// 读取本地数据
	var backup model.PanelAllEnv
	// 打开文件
	file, err := os.Open("./backup.json")
	if err != nil {
		// 打开文件时发生错误
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}
	// 延迟关闭
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			zap.L().Error(err.Error())
		}
	}(file)

	// 配置读取
	byteData, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		// 读取配置时发生错误
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	// 数据绑定
	err3 := json.Unmarshal(byteData, &backup.Data)
	if err3 != nil {
		// 数据绑定时发生错误
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	// 上传数据
	go EnvUp(backup, one.URL, one.Token, one.Params, "恢复任务")
	// 删除本地数据
	go DelBackupJSON()
	return res.CodeSuccess
}

// EnvUp 变量上传
func EnvUp(p model.PanelAllEnv, url, token string, params int, journal string) {
	var re int
	var panelData model.QLPanel

	for i := 0; i < len(p.Data); i++ {
		var data string
		var pRes model.PanelRes

		if re == 1 {
			panelData = sqlite.GetPanelDataByURL(url)
		}

		zap.L().Debug("URL地址：" + url)
		URL := panel.StringHTTP(url) + "/open/envs?t=" + strconv.Itoa(params)
		// 将字符串里面的双引号添加转义
		p.Data[i].Value = strings.ReplaceAll(p.Data[i].Value, `"`, `\"`)
		// 上传
		data = `[{"value": "` + p.Data[i].Value + `","name": "` + p.Data[i].Name + `","remarks": "` + p.Data[i].Remarks + `"}]`
		// 执行上传任务
		r, err := requests.Requests("POST", URL, data, token)
		if err != nil {
			// 记录错误
			zap.L().Error("[容器：变量：上传]请求发送失败:" + err.Error())
			sqlite.RecordingError(journal, err.Error())
		}
		// 序列化内容
		err = json.Unmarshal(r, &pRes)
		if err != nil {
			zap.L().Error(err.Error())
		}

		if pRes.Code >= 401 && pRes.Code <= 500 {
			// 更新Token, 再次提交
			_, t := panel.GetPanelToken(panel.StringHTTP(url), panelData.ClientID, panelData.ClientSecret)
			token = t.Data.Token
			params = t.Data.Expiration
			re = 1
			i -= 1
		} else if pRes.Code == 400 {
			// 可能是重复上传，跳过
			continue
		} else if pRes.Code >= 500 {
			// 青龙请求错误, 再次提交
			i -= 1
		} else if pRes.Code == 200 {
			re = 0
		}

		// 限速
		time.Sleep(time.Second / 8)
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

// GetConInfo 获取十条日志记录
func GetConInfo() ([]model.OperationRecord, res.ResCode) {
	// 查询记录
	info := sqlite.GetConData()
	return info, res.CodeSuccess
}

// DelBackupJSON 删除本地数据
func DelBackupJSON() {
	time.Sleep(time.Second * 10)
	err := os.Remove("backup.json")
	if err != nil {
		zap.L().Error(err.Error())
	}
}
