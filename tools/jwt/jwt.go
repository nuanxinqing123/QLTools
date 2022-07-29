// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 14:00
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : jwt.go

package jwt

import (
	"QLPanelTools/server/sqlite"
	"errors"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"time"
)

// TokenExpireDuration Token过期时间(7天)
const TokenExpireDuration = time.Hour * 24 * 7

// 加盐
//var mySecret = []byte("44tEeUAVxTKxcU6We9dc")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(userID int64, email string) (string, error) {
	// 获取密钥
	jwtKey := sqlite.GetJWTKey()

	// 加盐
	var mySecret = []byte(jwtKey)
	zap.L().Debug(jwtKey)

	// 创建声明数据
	c := MyClaims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "QLTools",                                  // 签发人
		},
	}

	// 使用指定的签名方式创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 使用指定的签名并获得完整编码后的Token
	return token.SignedString(mySecret)
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {
	// 获取密钥
	jwtKey := sqlite.GetJWTKey()

	// 加盐
	var mySecret = []byte(jwtKey)
	zap.L().Debug(jwtKey)

	// 解析Token
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	// 校验Token
	if token.Valid {
		return mc, nil
	}

	return nil, errors.New("invalid token")
}
