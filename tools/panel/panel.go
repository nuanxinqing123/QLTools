// -*- coding: utf-8 -*-
// @Time    : 2022/4/8 14:59
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : panel.go

package panel

import (
	"QLPanelTools/server/model"
	"QLPanelTools/server/sqlite"
	"QLPanelTools/tools/requests"
	res "QLPanelTools/tools/response"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

// StringHTTP 处理URL地址结尾的斜杠
func StringHTTP(url string) string {
	s := []byte(url)
	if s[len(s)-1] == '/' {
		s = s[:len(s)-1]
	}
	url = string(s)
	return url
}

// GetPanelToken 获取面板Token(有效期：30天)
func GetPanelToken(url, id, secret string) (res.ResCode, model.Token) {
	var token model.Token

	URL := StringHTTP(url) + fmt.Sprintf("/open/auth/token?client_id=%s&client_secret=%s", id, secret)

	// 请求Token
	strData, err := requests.Requests("GET", URL, "", "")
	if err != nil {
		return res.CodeServerBusy, token
	}

	// 序列化内容
	err = json.Unmarshal(strData, &token)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy, token
	}

	// 更新数据库储存Token
	go sqlite.SaveToken(url, token.Data.Token, token.Data.Expiration)

	return res.CodeSuccess, token
}

// TestGetPanelToken 测试面板连接
func TestGetPanelToken(url, id, secret string) (res.ResCode, model.Token) {
	var token model.Token

	URL := fmt.Sprintf("/open/auth/token?client_id=%s&client_secret=%s", id, secret)
	nUrl := StringHTTP(url)

	// 请求Token
	strData, err := requests.Requests("GET", nUrl+URL, "", "")
	if err != nil {
		return res.CodeServerBusy, token
	}

	// 序列化内容
	err = json.Unmarshal(strData, &token)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy, token
	}

	return res.CodeSuccess, token
}
