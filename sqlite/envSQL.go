// -*- coding: utf-8 -*-
// @Time    : 2022/4/6 17:15
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : envSQL.go

package sqlite

import (
	"QLPanelTools/model"
	"QLPanelTools/tools/timeTools"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

// CheckEnvName 检查变量名是否存在
func CheckEnvName(name string) bool {
	var env model.EnvName
	DB.Where("name = ?", name).First(env)
	if env.ID != 0 {
		// 变量名已存在
		return true
	} else {
		// 变量名不存在
		return false
	}
}

// AddEnvName 新增变量名
func AddEnvName(data *model.EnvNameAdd) (err error) {
	var e model.EnvName
	e.Name = data.EnvName
	e.NameRemarks = data.EnvNameRemarks
	e.Quantity = data.EnvQuantity
	e.Regex = data.EnvRegex
	e.Mode = data.EnvMode
	e.Division = data.EnvDivision
	e.ReUpdate = data.EnvReUpdate
	e.IsPlugin = data.EnvIsPlugin
	e.PluginName = data.EnvPluginName
	e.IsCDK = data.EnvIsCDK
	err = DB.Create(&e).Error
	if err != nil {
		zap.L().Error("Insert data error, err:", zap.Error(err))
		return
	}
	return
}

// UpdateEnvName 更新变量信息
func UpdateEnvName(data *model.EnvNameUp) {
	var d model.EnvName
	// 通过ID查询并更新数据
	DB.Where("id = ?", data.EnvID).First(&d)
	d.Name = data.EnvName
	d.NameRemarks = data.EnvNameRemarks
	d.Quantity = data.EnvQuantity
	d.Regex = data.EnvRegex
	d.Mode = data.EnvMode
	d.Division = data.EnvDivision
	d.ReUpdate = data.EnvReUpdate
	d.IsPlugin = data.EnvIsPlugin
	d.PluginName = data.EnvPluginName
	d.IsCDK = data.EnvIsCDK
	DB.Save(&d)
}

// DelEnvName 删除变量信息
func DelEnvName(data *model.EnvNameDel) {
	var d model.EnvName
	DB.Where("id = ? ", data.EnvID).First(&d)
	DB.Delete(&d)
}

// GetEnvNameAll 获取变量All数据
func GetEnvNameAll() []model.EnvName {
	var s []model.EnvName
	DB.Find(&s)
	return s
}

// GetEnvAllByID 根据ID值获取变量数据
func GetEnvAllByID(id int) []model.EnvName {
	// ID值查询服务
	var s model.QLPanel
	DB.First(&s, id)
	// 转换切片
	envBind := strings.Split(s.EnvBinding, "@")
	// 切片转换int类型
	var e []int
	zap.L().Debug("面板为：" + s.PanelName)
	for i := 0; i < len(envBind); i++ {
		zap.L().Debug("面板绑定变量数据为：" + envBind[i])
		if envBind[i] != "" {
			ee, err := strconv.Atoi(envBind[i])
			if err != nil {
				zap.L().Error(err.Error())
			}
			e = append(e, ee)
		}
	}

	var env []model.EnvName
	// 根据绑定值查询变量数据
	if len(e) != 0 {
		DB.Find(&env, e)
	}
	return env
}

// GetEnvNameCount 根据变量名获取配额
func GetEnvNameCount(name string) int {
	var env model.EnvName
	DB.Where("name = ?", name).First(&env)
	return env.Quantity
}

// CheckIPCount 查询IP今日已上传次数
func CheckIPCount(ip string, value string) bool {
	var p string
	todayTime := timeTools.SwitchTimeStampToDataYear(time.Now().Unix())
	sqlStr := "SELECT count(ip_address) FROM `ip_submit_records` WHERE submit_time = '" + todayTime + "' AND ip_address = '" + ip + "'"
	DB.Raw(sqlStr).Scan(&p)

	// 判断是否超出限制
	a, _ := strconv.Atoi(p)
	b, _ := strconv.Atoi(value)
	if a >= b {
		return true
	} else {
		return false
	}
}

// InsertSubmitRecord 记录上传IP
func InsertSubmitRecord(ip string) {
	var lr model.IPSubmitRecord
	lr.SubmitTime = timeTools.SwitchTimeStampToDataYear(time.Now().Unix())
	lr.IPAddress = ip
	DB.Create(&lr)
}
