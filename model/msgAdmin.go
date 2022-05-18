// -*- coding: utf-8 -*-
// @Time    : 2022/5/17 11:30
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : msgAdmin.go

package model

import "gorm.io/gorm"

type Email struct {
	gorm.Model
	EnableEmail bool   // 推送状态
	SendMail    string // 发件人邮箱
	SendPwd     string // 发件人密码
	SMTPServer  string // SMTP 邮件服务器地址
	SMTPPort    int    // SMTP端口
	SendName    string // 发件人昵称
}

type UpdateEmail struct {
	EnableEmail bool   `json:"enableEmail"`                   // 推送状态
	SendMail    string `json:"sendMail" binding:"required"`   // 发件人邮箱
	SendPwd     string `json:"sendPwd" binding:"required"`    // 发件人密码
	SMTPServer  string `json:"SMTPServer" binding:"required"` // SMTP 邮件服务器地址
	SMTPPort    int    `json:"SMTPPort" binding:"required"`   // SMTP端口
	SendName    string `json:"sendName" binding:"required"`   // 发件人昵称
}

type TestEmail struct {
	TEmail string `json:"TestEmail"` // 测试邮件发送
}
