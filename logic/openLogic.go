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
	"strings"
)

// EnvData 获取all data
func EnvData() (res.ResCode, model.EnvStartServer) {
	var sd model.EnvStartServer

	// 获取所有服务器信息
	sData := sqlite.GetServerCount()
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

	for i := 0; i < len(sd.ServerData); i++ {
		// 获取变量数据
		envData := sqlite.GetEnvAllByID(sd.ServerData[i].ID)
		if len(envData) != 0 {
			eData, err := json.Marshal(envData)
			if err != nil {
				zap.L().Error(err.Error())
				return res.CodeServerBusy, sd
			}

			// 数据绑定
			err = json.Unmarshal(eData, &sd.ServerData[i].EnvData)
			if err != nil {
				zap.L().Error(err.Error())
				return res.CodeServerBusy, sd
			}

			// 获取面板已存在变量数量
			url := panel.StringHTTP(sData[i].URL) + "/open/envs?searchValue=&t=" + strconv.Itoa(sData[i].Params)
			allData, err := requests.Requests("GET", url, "", sData[i].Token)
			if err != nil {
				return res.CodeServerBusy, sd
			}
			var token model.EnvData
			err = json.Unmarshal(allData, &token)
			if err != nil {
				zap.L().Error(err.Error())
				return res.CodeServerBusy, sd
			}

			// 判断返回状态
			if token.Code != 200 {
				// 尝试获取授权
				go panel.GetPanelToken(sData[i].URL, sData[i].ClientID, sData[i].ClientSecret)

				// 未授权或Token失效
				return res.CodeDataError, sd
			}

			// 计算变量剩余限额
			for x := 0; x < len(sd.ServerData[i].EnvData); x++ {
				sd.ServerData[i].EnvData[x].Quantity, _, _ = CalculateQuantity(sd.ServerData[i].ID, sd.ServerData[i].EnvData[x].Name)
			}
		}
	}
	return res.CodeSuccess, sd
}

// EnvAdd 添加变量
func EnvAdd(p *model.EnvAdd) res.ResCode {
	// 不允许内容为空
	if p.EnvData == "" {
		return res.CodeDataIsNull
	}

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

	// 转换切片
	envBind := strings.Split(sData.EnvBinding, "")
	// 校验变量是否处于容器白名单
	num := 0
	for i := 0; i < len(envBind); i++ {
		if envBind[i] == strconv.Itoa(int(eData.ID)) {
			num++
		}
	}
	if num == 0 {
		return res.CodeErrorOccurredInTheRequest
	}

	// 正则处理
	var s [][]string
	if eData.Regex != "" {
		// 需要处理正则
		zap.L().Debug("需要处理正则")
		reg := regexp.MustCompile(eData.Regex)
		// 匹配内容
		if reg != nil {
			s = reg.FindAllStringSubmatch(p.EnvData, -1)
			if len(s) == 0 {
				return res.CodeEnvDataMismatch
			}
		} else {
			return res.CodeServerBusy
		}
	}

	// 校验变量配额
	c, t, code := CalculateQuantity(p.ServerID, p.EnvName)
	if code == res.CodeServerBusy {
		zap.L().Debug("处理正则失败")
		return res.CodeServerBusy
	} else if c <= 0 {
		zap.L().Debug("限额已满，禁止提交")
		return res.CodeLocationFull
	}

	// 检查重复提交
	var bol bool
	var QCount int
	if eData.Regex != "" {
		bol, QCount = CheckRepeat(t, s[0][0], p.EnvName, eData)
	} else {
		bol, QCount = CheckRepeat(t, p.EnvData, p.EnvName, eData)
	}
	if bol == true {
		return res.CodeNoDuplicateSubmission
	}

	// 提交到服务器
	var data string
	url := panel.StringHTTP(sData.URL) + "/open/envs?t=" + strconv.Itoa(sData.Params)
	zap.L().Debug(url)
	if eData.Regex != "" {
		// 指定上传数据
		if eData.Mode == 1 {
			// 新建模式
			data = `[{"value": "` + s[0][0] + `","name": "` + p.EnvName + `","remarks": "` + p.EnvRemarks + `"}]`
		} else {
			// 合并模式
			if QCount != -1 {
				vv := t.Data[QCount].Value + eData.Division + s[0][0]
				p.EnvRemarks = t.Data[QCount].Name
				data = `{"id": "` + strconv.Itoa(t.Data[QCount].ID) + `", "value": "` + vv + `","name": "` + p.EnvName + `","remarks": "` + p.EnvRemarks + `"}`
			} else {
				data = `[{"value": "` + s[0][0] + `","name": "` + p.EnvName + `"}]`
			}
		}
	} else {
		// 指定上传数据
		if eData.Mode == 1 {
			// 新建模式
			data = `[{"value": "` + p.EnvData + `","name": "` + p.EnvName + `","remarks": "` + p.EnvRemarks + `"}]`
		} else {
			// 合并模式
			if QCount != -1 {
				vv := t.Data[QCount].Value + eData.Division + p.EnvData
				p.EnvRemarks = t.Data[QCount].Name
				data = `{"id": "` + strconv.Itoa(t.Data[QCount].ID) + `", "value": "` + vv + `","name": "` + p.EnvName + `","remarks": "` + p.EnvRemarks + `"}`
			} else {
				data = `[{"value": "` + p.EnvData + `","name": "` + p.EnvName + `"}]`
			}
		}
	}
	zap.L().Debug(data)
	var r []byte
	var err error
	if eData.Mode == 1 {
		// 新建模式(POST)
		r, err = requests.Requests("POST", url, data, sData.Token)
	} else {
		// 合并模式(PUT)
		zap.L().Debug(strconv.Itoa(QCount))
		if QCount != -1 {
			r, err = requests.Requests("PUT", url, data, sData.Token)
		} else {
			// 面板不存在合并模式变量时
			r, err = requests.Requests("POST", url, data, sData.Token)
		}
	}
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
func CalculateQuantity(id int, name string) (int, model.EnvData, res.ResCode) {
	var token model.EnvData
	// 获取变量数据
	count := sqlite.GetEnvNameCount(name)

	// 获取容器信息
	sData := sqlite.GetPanelDataByID(id)

	// 获取面板已存在变量数量
	url := panel.StringHTTP(sData.URL) + "/open/envs?searchValue=&t=" + strconv.Itoa(sData.Params)
	allData, err := requests.Requests("GET", url, "", sData.Token)
	if err != nil {
		return 0, token, res.CodeServerBusy
	}

	err = json.Unmarshal(allData, &token)
	if err != nil {
		zap.L().Error(err.Error())
		return 0, token, res.CodeServerBusy
	}

	// 计算变量剩余限额
	c := count
	for i := 0; i < len(token.Data); i++ {
		if token.Data[i].Name == name {
			c--
		}
	}

	return c, token, res.CodeSuccess
}

// CheckRepeat 校验是否重复上传
func CheckRepeat(p model.EnvData, env, name string, data model.EnvName) (bool, int) {
	var QCount = -1
	// 通过变量名获取上传模式
	if data.Mode == 1 {
		// 新建模式
		var count = 0
		for i := 0; i < len(p.Data); i++ {
			if p.Data[i].Value == env {
				count++
				break
			}
		}
		if count != 0 {
			return true, 0
		}
	} else {
		// 合并模式
		var count = 0
		// 遍历所有表获取合并表
		if len(p.Data) == 0 {
			return false, QCount
		}
		for i := 0; i < len(p.Data); i++ {
			if p.Data[i].Name == name {
				count = i
				QCount = i
				break
			}
		}

		// 根据分隔符处理面板上的数据
		var up = 0
		envList := strings.Split(p.Data[count].Value, data.Division)
		for i := 0; i < len(envList); i++ {
			if envList[i] == env {
				up++
				break
			}
		}
		if up != 0 {
			return true, 0
		}
	}
	return false, QCount
}
