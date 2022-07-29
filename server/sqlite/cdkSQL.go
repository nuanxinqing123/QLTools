// -*- coding: utf-8 -*-
// @Time    : 2022/6/12 14:37
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : cdkSQL.go

package sqlite

import (
	"QLPanelTools/server/model"
	"go.uber.org/zap"
	"strconv"
)

// GetDivisionCDKData 条件查询CDK数据
func GetDivisionCDKData(page int) []model.CDK {
	var cdk []model.CDK
	if page == 1 {
		// 获取前20条数据
		sql := "select * from cdks where `deleted_at` IS NULL limit 0, " + strconv.Itoa(19) + ";"
		zap.L().Debug(sql)
		DB.Raw(sql).Scan(&cdk)
	} else {
		/*
			1、获取指定页数的数据
			2、20条数据为一页
			3、起始位置：((page - 1) * 20) - 1
			4、结束位置：(page * 20) - 1
		*/
		sql := "select * from cdks where `deleted_at` IS NULL limit " + strconv.Itoa(((page-1)*20)-1) + ", " + strconv.Itoa((page*20)-1) + ";"
		zap.L().Debug(sql)
		DB.Raw(sql).Scan(&cdk)
	}
	return cdk
}

// GetCDKDataPage 获取CDK表总数据
func GetCDKDataPage() int64 {
	var c []model.CDK
	result := DB.Find(&c)
	return result.RowsAffected
}

// GetAllCDKData 查询全部CDK数据
func GetAllCDKData() []model.CDK {
	var c []model.CDK
	DB.Find(&c)
	return c
}

// GetTrueCDKData 查询启用CDK数据
func GetTrueCDKData() []model.CDK {
	var c []model.CDK
	DB.Where("state = ?", true).Find(&c)
	return c
}

// GetFalseCDKData 查询禁用CDK数据
func GetFalseCDKData() []model.CDK {
	var c []model.CDK
	DB.Where("state = ?", false).Find(&c)
	return c
}

// GetOneCDKData 查询指定CDK数据
func GetOneCDKData(cdk string) []model.CDK {
	var c []model.CDK
	DB.Where("cd_key = ?", cdk).First(&c)
	return c
}

// InsertCDKData 生成CDK写入数据库
func InsertCDKData(p *model.CDK) {
	var cdk model.CDK
	cdk.CdKey = p.CdKey
	cdk.AvailableTimes = p.AvailableTimes
	cdk.State = p.State
	DB.Create(&cdk)
}

// UpdateCDKDataState 批量更新CDK状态
func UpdateCDKDataState(p *model.UpdateStateCDK) {
	cdk := new(model.CDK)
	if p.State == 1 {
		// 启用
		zap.L().Debug("CDK全部：启用")
		DB.Model(&cdk).Where("state = ?", false).Update("state", true)
	} else {
		// 禁用
		zap.L().Debug("CDK全部：禁用")
		DB.Model(&cdk).Where("state = ?", true).Update("state", false)
	}
}

// UpdateCDKData 更新单条CDK数据
func UpdateCDKData(p *model.UpdateCDK) {
	cdk := new(model.CDK)
	DB.Where("id = ?", p.ID).First(&cdk)
	cdk.AvailableTimes = p.AvailableTimes
	cdk.State = p.State
	DB.Save(&cdk)
}

// DelCDKData 删除CDK数据
func DelCDKData(p *model.DelCDK) {
	cdk := new(model.CDK)
	DB.Where("id = ? ", p.ID).First(&cdk)
	DB.Delete(&cdk)
}

// GetCDKData 查询CDK信息
func GetCDKData(p string) model.CDK {
	var c model.CDK
	DB.Where("cd_key = ?", p).First(&c)
	return c
}
