// -*- coding: utf-8 -*-
// @Time    : 2022/5/17 12:30
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : messageLogic.go

package logic

import (
	"QLPanelTools/server/model"
	"QLPanelTools/server/sqlite"
	"QLPanelTools/tools/email"
	res "QLPanelTools/tools/response"
	"go.uber.org/zap"
)

// GetEmailData 获取邮件信息
func GetEmailData() (res.ResCode, model.Email) {
	return res.CodeSuccess, sqlite.GetEmailOne()
}

// TestEmailSend 测试服务发送
func TestEmailSend(p *model.TestEmail) res.ResCode {
	// 定义收信人
	mailTo := []string{p.TEmail}
	err := email.SendMail(mailTo, "青龙Tools", "青龙Tools邮件测试发送")
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeServerBusy
	}

	return res.CodeSuccess
}

// UpdateEmailSet 更新邮件服务信息
func UpdateEmailSet(p *model.UpdateEmail) res.ResCode {
	// 获取第一条配置信息
	emailData := sqlite.GetEmailOne()
	if emailData.SendMail == "" && emailData.SendPwd == "" && emailData.SMTPServer == "" {
		// 首次修改, 创建
		emailData.EnableEmail = p.EnableEmail
		emailData.SendName = p.SendName
		emailData.SendMail = p.SendMail
		emailData.SendPwd = p.SendPwd
		emailData.SMTPServer = p.SMTPServer
		emailData.SMTPPort = p.SMTPPort
		emailData.SendName = p.SendName
		sqlite.InsertEmailData(emailData)
	} else {
		sqlite.UpdateEmailData(p)
	}

	return res.CodeSuccess
}
