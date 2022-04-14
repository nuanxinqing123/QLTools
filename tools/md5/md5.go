// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 14:02
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : md5.go

package md5

import (
	"crypto/md5"
	"fmt"
)

// AddMD5 MD5加密
func AddMD5(text string) string {
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(text)))
	return md5Str
}
