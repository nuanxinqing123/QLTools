// -*- coding: utf-8 -*-
// @Time    : 2022/6/12 14:33
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : cdkLogic.go

package logic

import (
	"QLPanelTools/model"
	"QLPanelTools/sqlite"
	res "QLPanelTools/tools/response"
	"bufio"
	"github.com/segmentio/ksuid"
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"
)

// GetDivisionCDKData CDK：以20条数据分割
func GetDivisionCDKData(page string) (res.ResCode, model.CDKPageData) {
	var data []model.CDK
	var cdkPage model.CDKPageData
	if page == "" {
		// 空值，默认获取前20条数据(第一页)
		data = sqlite.GetDivisionCDKData(1)
	} else {
		// String转Int
		intPage, err := strconv.Atoi(page)
		if err != nil {
			// 类型转换失败，查询默认获取前20条数据(第一页)
			zap.L().Error(err.Error())
			data = sqlite.GetDivisionCDKData(1)
		} else {
			// 查询指定页数的数据
			data = sqlite.GetDivisionCDKData(intPage)
		}
	}

	// 查询总页数
	count := sqlite.GetCDKDataPage()
	// 计算页数
	z := count / 20
	var y int64
	y = count % 20

	if y != 0 {
		cdkPage.Page = z + 1
	} else {
		cdkPage.Page = z
	}
	cdkPage.CDKData = data

	return res.CodeSuccess, cdkPage
}

// GetCDKData 条件查询CDK数据
func GetCDKData(state int) (res.ResCode, []model.CDK) {
	// state 1：全部、2：启用、3：禁用
	var data []model.CDK
	if state == 1 {
		data = sqlite.GetAllCDKData()
	} else if state == 2 {
		data = sqlite.GetTrueCDKData()
	} else {
		data = sqlite.GetFalseCDKData()
	}

	return res.CodeSuccess, data
}

// GetOneCDKData 搜索CDK数据
func GetOneCDKData(cdk string) (res.ResCode, []model.CDK) {
	return res.CodeSuccess, sqlite.GetOneCDKData(cdk)
}

// CreateCDKData 生成CDK数据
func CreateCDKData(p *model.CreateCDK) res.ResCode {
	// 判断本地是否还有遗留文件
	_, err := os.Stat("CDK.txt")
	if err == nil {
		// 删除旧文件
		err := os.Remove("CDK.txt")
		if err != nil {
			zap.L().Error(err.Error())
			return res.CodeServerBusy
		}
	}

	// 创建记录数组
	var li []string

	// 获取生成数量
	for i := 0; i < p.CdKeyCount; i++ {
		// 创建对象
		cdk := new(model.CDK)

		// 生成用户UID
		uid := ksuid.New()
		cdk.CdKey = uid.String()
		cdk.AvailableTimes = p.CdKeyCount
		cdk.State = true

		// 加入数组
		li = append(li, cdk.CdKey)

		// 写入数据库
		sqlite.InsertCDKData(cdk)
	}

	// 创建CDK.txt并写入数据
	filepath := "CDK.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0)
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeCreateFileError
	}
	defer file.Close()

	// 写入CDK数据
	writer := bufio.NewWriter(file)
	for i := 0; i < len(li); i++ {
		_, err2 := writer.WriteString(li[i] + "\n")
		if err2 != nil {
			zap.L().Error(err2.Error())
			return res.CodeCreateFileError
		}
	}
	err = writer.Flush()
	if err != nil {
		zap.L().Error(err.Error())
		return res.CodeCreateFileError
	}

	return res.CodeSuccess
}

// DelCDKData 删除本地数据
func DelCDKData() {
	time.Sleep(time.Second * 10)
	err := os.Remove("CDK.txt")
	if err != nil {
		zap.L().Error(err.Error())
	}
}

// CDKState CDK全部启用/禁用
func CDKState(p *model.UpdateStateCDK) res.ResCode {
	sqlite.UpdateCDKDataState(p)
	return res.CodeSuccess
}

// UpdateCDK 更新单条CDK数据
func UpdateCDK(p *model.UpdateCDK) res.ResCode {
	sqlite.UpdateCDKData(p)
	return res.CodeSuccess
}

// DelCDK 删除CDK
func DelCDK(p *model.DelCDK) res.ResCode {
	sqlite.DelCDKData(p)
	return res.CodeSuccess
}
