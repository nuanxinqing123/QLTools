// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 14:21
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : userSQL.go

package sqlite

import (
	"QLPanelTools/server/model"
	res "QLPanelTools/tools/response"
	"QLPanelTools/tools/timeTools"
	"time"
)

// InsertUser 创建新用户
func InsertUser(user *model.User) (err error) {
	err = DB.Create(&user).Error
	if err != nil {
		return
	}
	return
}

// GetUserData 获取用户信息
func GetUserData() (bool, model.User) {
	var user model.User
	// 获取键值第一条记录
	DB.First(&user)

	// 判断是否已注册
	if user.Username != "" {
		return true, user
	} else {
		return false, user
	}
}

// CheckEmail 检查邮箱是否存在
func CheckEmail(email string) res.ResCode {
	var user model.User
	DB.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		// 邮箱已存在
		return res.CodeEmailExist
	} else {
		// 邮箱不存在
		return res.CodeEmailNotExist
	}
}

// CheckAdmin 判断是否属于管理员
func CheckAdmin(userID interface{}) bool {
	var user model.User
	DB.Where("user_id = ?", userID).First(&user)
	if user.Username != "" {
		return true
	} else {
		return false
	}
}

// UpdateUserData 更新用户信息
func UpdateUserData(email, pwd string) error {
	var user model.User
	// 获取管理员信息
	result := DB.First(&user)
	if result.Error != nil {
		return result.Error
	}
	// 更新数据
	if email != "" {
		user.Email = email
	}
	user.Password = pwd
	// 储存数据
	result = DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// InsertLoginRecord 创建登录记录
func InsertLoginRecord(lr *model.LoginRecord) {
	DB.Create(&lr)
}

// GetIPData 获取IP记录
func GetIPData() []model.IpData {
	var i []model.IpData
	sqlStr := "SELECT `login_time`, `ip`, `ip_address`, `if_ok` FROM `login_records` order by id desc limit 0,10;"
	DB.Raw(sqlStr).Scan(&i)
	return i
}

// GetFailLoginIPData 获取当日登录IP记录
func GetFailLoginIPData() []model.IpData {
	var i []model.IpData
	sqlStr := "SELECT `login_day`, `login_time`, `ip`, `ip_address`, `if_ok` FROM `login_records` WHERE login_day = '" + timeTools.SwitchTimeStampToDataYear(time.Now().Unix()) + "';"
	DB.Raw(sqlStr).Scan(&i)
	return i
}
