// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 14:18
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : userLogic.go

package logic

import (
	"QLPanelTools/model"
	"QLPanelTools/sqlite"
	"QLPanelTools/tools/email"
	"QLPanelTools/tools/jwt"
	"QLPanelTools/tools/md5"
	"QLPanelTools/tools/requests"
	res "QLPanelTools/tools/response"
	"QLPanelTools/tools/snowflake"
	"QLPanelTools/tools/timeTools"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// SignUp 注册业务
func SignUp(p *model.UserSignUp) res.ResCode {
	// 判断是否已存在账户
	result, _ := sqlite.GetUserData()
	if result == true {
		return res.CodeRegistrationClosed
	}

	// 密码加密
	p.Password = md5.AddMD5(p.Password)

	// 生成用户UID
	userID := snowflake.GenID()

	// 构造User实例
	user := &model.User{
		UserID:   userID,
		Email:    p.Email,
		Username: p.Username,
		Password: p.Password,
	}

	// 保存进数据库
	err := sqlite.InsertUser(user)
	if err != nil {
		zap.L().Error("Error inserting database, err:", zap.Error(err))
		return res.CodeServerBusy
	}
	return res.CodeSuccess
}

// SignIn 登录业务
func SignIn(p *model.UserSignIn) (string, res.ResCode) {
	// 检查邮箱格式是否正确
	bol := email.VerifyEmailFormat(p.Email)
	if bol == false {
		return "", res.CodeEmailFormatError
	} else {
		// 检查邮箱是否存在
		code := sqlite.CheckEmail(p.Email)
		if code == res.CodeEmailNotExist {
			// 不存在
			return "", res.CodeEmailNotExist
		} else {
			// 邮箱存在,记录传入密码
			oPassword := p.Password

			// 获取数据库用户信息
			_, user := sqlite.GetUserData()

			// 判断密码是否正确
			if user.Password != md5.AddMD5(oPassword) {
				return "", res.CodeInvalidPassword
			} else {
				// 密码正确, 返回生成的Token
				token, err := jwt.GenToken(user.UserID, user.Email)
				if err != nil {
					zap.L().Error("An error occurred in token generation, err:", zap.Error(err))
					return "", res.CodeServerBusy
				}
				return token, res.CodeSuccess
			}
		}
	}
}

// RePwd 修改密码业务
func RePwd(p *model.ReAdminPwd) (bool, res.ResCode) {
	// 获取数据库用户信息
	_, user := sqlite.GetUserData()

	// 判断密码是否正确
	if user.Password != md5.AddMD5(p.OldPassword) {
		return false, res.CodeOldPassWordError
	} else {
		// 储存新密码
		err := sqlite.UpdateUserData(md5.AddMD5(p.Password))
		if err != nil {
			return false, res.CodeServerBusy
		}
		return true, res.CodeSuccess
	}

}

// CheckToken 检查Token是否有效
func CheckToken(p *model.CheckToken) (bool, res.ResCode) {
	// 获取管理员信息
	_, aData := sqlite.GetUserData()

	// 解析Token
	myClaims, err := jwt.ParseToken(p.JWToken)
	if err != nil {
		return false, res.CodeServerBusy
	}

	if aData.Email != myClaims.Email {
		return false, res.CodeInvalidToken
	}
	if aData.UserID != myClaims.UserID {
		return false, res.CodeInvalidToken
	}

	zap.L().Debug(strconv.FormatInt(myClaims.UserID, 10))

	return true, res.CodeSuccess
}

// AddIPAddr 记录登录信息
func AddIPAddr(ip string, ifok bool) {
	// 查询IP地址
	url := "https://api.iyk0.com/ip/?ip=" + ip
	addr, err := requests.Requests("GET", url, "", "")
	if err != nil {
		return
	}

	ipCreate := &model.LoginRecord{
		LoginTime: timeTools.SwitchTimeStampToData(time.Now().Unix()),
		IP:        ip,
		IPAddress: fmt.Sprintf("%s", addr),
		IfOK:      ifok,
	}
	// 储存记录
	sqlite.InsertLoginRecord(ipCreate)
}

// GetIPInfo 查询近十条记录
func GetIPInfo() ([]model.IpData, res.ResCode) {
	// 查询记录
	ip := sqlite.GetIPData()
	return ip, res.CodeSuccess
}

// GetAdminInfo 获取管理员信息
func GetAdminInfo() (string, res.ResCode) {
	_, info := sqlite.GetUserData()
	return info.Username, res.CodeSuccess
}
