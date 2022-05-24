// -*- coding: utf-8 -*-
// @Time    : 2022/4/23 15:33
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : systemLogic.go

package logic

import (
	_const "QLPanelTools/const"
	"QLPanelTools/model"
	"QLPanelTools/sqlite"
	"QLPanelTools/tools/requests"
	res "QLPanelTools/tools/response"
	"encoding/json"
	"github.com/staktrace/go-update"
	"go.uber.org/zap"
	"runtime"
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
	w.Version = _const.Version

	return w, res.CodeSuccess
}

// UpdateSoftware 更新软件
func UpdateSoftware(p *model.SoftWareGOOS) (res.ResCode, string) {
	if runtime.GOOS == "windows" {
		return res.CodeUpdateServerBusy, "Windows系统不支持此功能"
	}
	// 获取版本号
	var v model.Ver
	url := "https://version.6b7.xyz/qltools_version.json"
	r, _ := requests.Requests("GET", url, "", "")
	_ = json.Unmarshal(r, &v)
	if v.Version == _const.Version {
		return res.CodeUpdateServerBusy, "已经是最新版本"
	}

	// 更新程序
	go UpdateSoftWare(v.Version, p.Framework)

	return res.CodeSuccess, "程序已进入自动更新任务，如果更新失败请手动更新"
}

func UpdateSoftWare(version, GOOS string) {
	// 获取代理下载地址
	gh, _ := sqlite.GetSetting("ghProxy")

	var url string
	url = AddStringHTTP(gh.Value) + "https://github.com/nuanxinqing123/QLTools/releases/download/" + version
	if GOOS == "amd64" {
		url += "/QLTools-linux-amd64"
	} else if GOOS == "arm64" {
		url += "/QLTools-linux-arm64"
	} else {
		url += "/QLTools-linux-arm"
	}
	zap.L().Debug("Download: " + url)

	err := doUpdate(url)
	if err != nil {
		zap.L().Error(err.Error())
	}
}

func doUpdate(url string) error {
	resp, err := requests.Down(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		// error handling
		return err
	}
	return nil
}

// AddStringHTTP 处理URL地址结尾的斜杠
func AddStringHTTP(url string) string {
	if len(url) == 0 {
		return url
	}
	s := []byte(url)
	if s[len(s)-1] != '/' {
		url += "/"
	}
	return url
}
