// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 14:57
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : email.go

package email

import (
	"QLPanelTools/server/sqlite"
	"gopkg.in/gomail.v2"
	"regexp"
)

// VerifyEmailFormat 正则验证邮箱格式
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// SendMail 发送邮件
func SendMail(mailTo []string, subject string, body string) error {
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码
	// 获取邮件服务器信息
	es := sqlite.GetEmailOne()

	m := gomail.NewMessage()

	//这种方式可以添加别名，即“XX官方”
	m.SetHeader("From", m.FormatAddress(es.SendMail, es.SendName))
	// 说明：如果是用网易邮箱账号发送，以下方法别名可以是中文
	// 如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	// 发送给多个用户
	m.SetHeader("To", mailTo...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)

	d := gomail.NewDialer(es.SMTPServer, es.SMTPPort, es.SendMail, es.SendPwd)

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
