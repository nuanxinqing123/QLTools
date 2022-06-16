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
	"math/rand"
	"time"
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
		&model.IPSubmitRecord{},
		&model.Email{},
		&model.CDK{},
		&model.JWTAdmin{},
	)
	if err != nil {
		zap.L().Error("SQLite 自动迁移失败")
		panic(err.Error())
	}

	return DB
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func InitWebSettings() {
	// 判断Settings是否是第一次创建
	settings, err := GetSettings()
	if err != nil {
		zap.L().Error("InitWebSettings 发生错误")
		panic(err.Error())
	}

	// 检查JWT密钥表是否存在
	jwtKey := GetJWTKey()
	if jwtKey == "" || len(jwtKey) < 10 {
		// 生成密码并写入数据库
		b := make([]rune, 18)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := range b {
			b[i] = letters[r.Intn(62)]
		}
		zap.L().Debug("生成密钥：" + string(b))
		CreateJWTKey(string(b))
	}

	if len(settings) == 0 {
		zap.L().Debug("Init WebSettings")
		p := &[]model.WebSettings{
			{Key: "notice", Value: ""},
			{Key: "blacklist", Value: ""},
			{Key: "backgroundImage", Value: ""},
			{Key: "ipCount", Value: "0"},
			{Key: "ghProxy", Value: "https://ghproxy.com"},
			{Key: "webTitle", Value: "青龙Tools"},
		}

		err = SaveSettings(p)
		if err != nil {
			zap.L().Error("InitWebSettings 发生错误")
			panic(err.Error())
		}
	}
}
