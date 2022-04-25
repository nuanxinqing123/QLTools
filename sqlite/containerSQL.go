// -*- coding: utf-8 -*-
// @Time    : 2022/4/25 18:20
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : containerSQL.go

package sqlite

import "QLPanelTools/model"

// RecordingError 记录错误
func RecordingError(journal, info string) {
	var er model.OperationRecord
	// 记录日志
	er.Journal = info
	// 操作方式
	er.Operation = journal
	DB.Create(&er)
}
