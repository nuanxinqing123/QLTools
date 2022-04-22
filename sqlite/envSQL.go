// -*- coding: utf-8 -*-
// @Time    : 2022/4/6 17:15
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : envSQL.go

package sqlite

import (
	"QLPanelTools/model"
	"go.uber.org/zap"
	"strconv"
	"strings"
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
	e.Quantity = data.EnvQuantity
	e.Regex = data.EnvRegex
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
	d.Quantity = data.EnvQuantity
	d.Regex = data.EnvRegex
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
	envBind := strings.Split(s.EnvBinding, "")
	// 切片转换int类型
	var e []int
	for i := 0; i < len(envBind); i++ {
		ee, err := strconv.Atoi(envBind[i])
		if err != nil {
			zap.L().Error(err.Error())
		}
		e = append(e, ee)
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
