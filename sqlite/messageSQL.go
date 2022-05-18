// -*- coding: utf-8 -*-
// @Time    : 2022/5/17 12:32
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : messageSQL.go

package sqlite

import "QLPanelTools/model"

func InsertEmailData(email model.Email) {
	DB.Create(&email)
}

func UpdateEmailData(email *model.UpdateEmail) {
	var emailData model.Email
	DB.First(&emailData)
	emailData.EnableEmail = email.EnableEmail
	emailData.SendMail = email.SendMail
	emailData.SendPwd = email.SendPwd
	emailData.SMTPServer = email.SMTPServer
	emailData.SMTPPort = email.SMTPPort
	emailData.SendName = email.SendName
	DB.Save(&emailData)
}

func GetEmailOne() model.Email {
	var emailData model.Email
	DB.First(&emailData)
	return emailData
}
