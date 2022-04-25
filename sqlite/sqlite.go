// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 13:08
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : sqlite.go

package sqlite

import (
	"QLPanelTools/model"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() *gorm.DB {
	// 连接MySQL
	DB, err = gorm.Open(sqlite.Open("config/app.db"), &gorm.Config{})
	if err != nil {
		zap.L().Error("SQLite 发生错误")
		panic(err.Error())
	}

	// 自动迁移
	err := DB.AutoMigrate(
		&model.User{},
		&model.EnvName{},
		&model.QLPanel{},
		&model.LoginRecord{},
		&model.WebSettings{},
		&model.OperationRecord{},
	)
	if err != nil {
		zap.L().Error("SQLite 自动迁移失败")
		panic(err.Error())
	}

	return DB
}
