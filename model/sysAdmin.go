// -*- coding: utf-8 -*-
// @Time    : 2022/4/23 16:00
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : sysAdmin.go

package model

type Ver struct {
	Version string `json:"Version"`
	Notice  string `json:"Notice"`
}

type WebVer struct {
	Update  bool   `json:"Update"`
	Version string `json:"Version"`
	Notice  string `json:"Notice"`
}
