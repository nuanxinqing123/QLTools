// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 14:03
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : code.go

package response

type ResCode int64

const (
	CodeSuccess ResCode = 2000 + iota
	CodeEnvIsNull

	CodeInvalidParam = 5000 + iota
	CodeServerBusy
	CodeInvalidRouterRequested

	CodeInvalidToken
	CodeNeedLogin
	CodeRegistrationClosed
	CodeInvalidPassword
	CodeEmailFormatError
	CodeEmailNotExist
	CodeEmailExist

	CodeEnvNameExist

	CodeConnectionTimedOut
	CodeDataError
	CodeErrorOccurredInTheRequest
	CodeStorageFailed

	CodeCheckDataNotExist
	CodeOldPassWordError
	CodeEnvDataMismatch
	CodeLocationFull
	CodeDataIsNull
	CodePanelNotWhitelisted
	CodeNotStandardized
	CodeNoDuplicateSubmission
	CodeBlackListEnv
	CodeNumberDepletion
	CodeRemoveFail
	CodeCustomError
	CodeNoAdmittance
	CodeUpdateServerBusy
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:   "Success",
	CodeEnvIsNull: "当前面板没有变量,请直接添加",

	CodeInvalidParam:           "请求参数错误",
	CodeServerBusy:             "服务繁忙",
	CodeInvalidRouterRequested: "请求无效路由",

	CodeInvalidToken:       "无效的Token",
	CodeNeedLogin:          "未登录",
	CodeRegistrationClosed: "已关闭注册",
	CodeInvalidPassword:    "邮箱或密码错误",
	CodeEmailFormatError:   "邮箱格式错误",
	CodeEmailNotExist:      "邮箱不存在",
	CodeEmailExist:         "邮箱已存在",

	CodeEnvNameExist: "变量名已存在",

	CodeConnectionTimedOut:        "面板地址连接超时",
	CodeDataError:                 "面板信息有错误",
	CodeErrorOccurredInTheRequest: "请求发生错误",
	CodeStorageFailed:             "发生一点小意外，请重新提交",

	CodeCheckDataNotExist:     "查询信息为空",
	CodeOldPassWordError:      "旧密码错误",
	CodeEnvDataMismatch:       "上传内容不符合规定",
	CodeLocationFull:          "限额已满，禁止提交",
	CodeDataIsNull:            "提交内容不能为空",
	CodePanelNotWhitelisted:   "提交服务器不在白名单",
	CodeNotStandardized:       "提交数据不规范",
	CodeNoDuplicateSubmission: "禁止提交重复数据",
	CodeBlackListEnv:          "存在于黑名单的变量",
	CodeNumberDepletion:       "次数耗尽",
	CodeRemoveFail:            "删除失败",
	CodeCustomError:           "自定义错误",
	CodeNoAdmittance:          "数据禁止通过",
	CodeUpdateServerBusy:      "升级错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
