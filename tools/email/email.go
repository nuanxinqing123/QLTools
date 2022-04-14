// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 14:57
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : email.go

package email

import "regexp"

// VerifyEmailFormat 正则验证邮箱格式
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
