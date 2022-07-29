// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 14:05
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : auth.go

package middleware

import (
	"QLPanelTools/server/sqlite"
	"QLPanelTools/tools/jwt"
	res "QLPanelTools/tools/response"
	"strings"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

// UserAuth 基于JWT的认证中间件
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Token放在Header的Authorization中，并使用Bearer开头
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			res.ResError(c, res.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			res.ResError(c, res.CodeInvalidToken)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			res.ResError(c, res.CodeInvalidToken)
			c.Abort()
			return
		}

		// 检查是否属于管理员
		cAdmin := sqlite.CheckAdmin(mc.UserID)

		if cAdmin != true {
			c.Abort()
			res.ResErrorWithMsg(c, res.CodeInvalidRouterRequested, "无访问权限或认证已过期")
			return
		} else {
			//将当前请求的userID信息保存到请求的上下文c上
			c.Set(CtxUserIDKey, mc.UserID)
			c.Next()
		}
	}
}
