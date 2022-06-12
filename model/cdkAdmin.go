// -*- coding: utf-8 -*-
// @Time    : 2022/6/12 14:09
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : cdkAdmin.go

package model

import "gorm.io/gorm"

type CDK struct {
	gorm.Model
	CdKey          string // CD-KEY值
	AvailableTimes int    // CD-KEY剩余可用次数
	State          bool   // CD-KEY状态（true：启用、false：禁用）
}

type CreateCDK struct {
	CdKeyCount          int `json:"cdKeyCount" binding:"required"`          // CDK生成数量
	CdKeyAvailableTimes int `json:"cdKeyAvailableTimes" binding:"required"` // CDK使用次数
}

type UpdateStateCDK struct {
	State int `json:"state" binding:"required"` // CKD状态（1:启用、2:禁用）
}

type UpdateCDK struct {
	ID             int  `json:"id" binding:"required"` // CDK ID
	AvailableTimes int  `json:"availableTimes"`        // CDK使用次数
	State          bool `json:"state"`                 // CKD状态
}

type DelCDK struct {
	ID int `json:"id" binding:"required"` // CDK ID
}

type CDKPageData struct {
	Page    int64 `json:"page"`    // 总页数
	CDKData []CDK `json:"CDKData"` // 分页查询CDK数据
}
