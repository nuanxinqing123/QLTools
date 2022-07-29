// -*- coding: utf-8 -*-
// @Time    : 2022/6/16 19:15
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : jwtSQL.go

package sqlite

import "QLPanelTools/server/model"

func GetJWTKey() string {
	var jwt model.JWTAdmin
	DB.First(&jwt)
	return jwt.SecretKey
}

func CreateJWTKey(pwd string) {
	var jwt model.JWTAdmin
	jwt.SecretKey = pwd
	DB.Create(&jwt)
}
