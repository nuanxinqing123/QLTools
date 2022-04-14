// -*- coding: utf-8 -*-
// @Time    : 2022/4/7 16:30
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : tokenSQL.go

package sqlite

import (
	"QLPanelTools/model"
	"go.uber.org/zap"
)

// SaveToken 储存Token
func SaveToken(url, token string, params int) {
	var t model.QLPanel
	zap.L().Debug("SaveTokenUrl:" + url)
	// 通过URL查询并更新数据
	DB.Where("url = ?", url).First(&t)
	t.Token = token
	t.Params = params
	DB.Save(&t)
}
