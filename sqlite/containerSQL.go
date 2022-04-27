// -*- coding: utf-8 -*-
// @Time    : 2022/4/25 18:20
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : containerSQL.go

package sqlite

import (
	"QLPanelTools/model"
	"QLPanelTools/tools/timeTools"
	"time"
)

// RecordingError 记录错误
func RecordingError(journal, info string) {
	var er model.OperationRecord
	// 发生时间
	er.OccurrenceTime = timeTools.SwitchTimeStampToData(time.Now().Unix())
	// 记录日志
	er.Journal = info
	// 操作方式
	er.Operation = journal
	DB.Create(&er)
}

// GetConData 获取十条错误记录
func GetConData() []model.OperationRecord {
	var i []model.OperationRecord
	sqlStr := "SELECT `occurrence_time`, `journal`, `operation` FROM `operation_records` order by id desc limit 0,10;"
	DB.Raw(sqlStr).Scan(&i)
	return i
}
