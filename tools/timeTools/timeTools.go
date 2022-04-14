// -*- coding: utf-8 -*-
// @Time    : 2022/4/12 20:14
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : timeTools.go

package timeTools

import "time"

// SwitchTimeStampToData 将传入的时间戳转为时间
func SwitchTimeStampToData(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
