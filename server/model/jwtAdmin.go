// -*- coding: utf-8 -*-
// @Time    : 2022/6/16 19:13
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : jwtAdmin.go

package model

import "gorm.io/gorm"

// JWTAdmin JWT密钥
type JWTAdmin struct {
	gorm.Model
	SecretKey string
}
