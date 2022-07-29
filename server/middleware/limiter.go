// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 14:46
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : limiter.go

package middleware

import (
	res "QLPanelTools/tools/response"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

// RateLimitMiddleware 限流
func RateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			res.ResErrorWithMsg(c, res.CodeServerBusy, "异常流量")
			c.Abort()
			return
		}
		c.Next()
	}
}
