// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 14:07
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : userAdmin.go

package model

import (
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	/*
		gorm.Model：基础结构（ID、CreatedAt、UpdatedAt、DeletedAt）
	*/
	gorm.Model
	// 用户ID
	UserID int64 `binding:"required"`
	// 用户邮箱
	Email string `binding:"required"`
	// 用户名
	Username string `binding:"required"`
	// 用户密码
	Password string `binding:"required"`
}

// LoginRecord 登录记录模型
type LoginRecord struct {
	gorm.Model
	LoginTime string `binding:"required"`
	IP        string `binding:"required"`
	IPAddress string `binding:"required"`
	IfOK      bool   `binding:"required"`
}

// UserSignUp 用户注册模型
type UserSignUp struct {
	Email      string `json:"email" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// UserSignIn 用户登录模型
type UserSignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ReAdminPwd 修改密码模型
type ReAdminPwd struct {
	OldPassword string `json:"old_password" binding:"required"`
	Password    string `json:"password" binding:"required"`
	RePassword  string `json:"re_password" binding:"required,eqfield=Password"`
}

// CheckToken 检查Token是否有效
type CheckToken struct {
	JWToken string `json:"token" binding:"required"`
}

// IPModel IP模型
type IPModel struct {
	Data []IpData `json:"Data"`
}

type IpData struct {
	LoginTime string `json:"login_time"`
	IP        string `json:"ip"`
	IPAddress string `json:"ip_address"`
	IfOK      bool   `json:"if_ok"`
}
