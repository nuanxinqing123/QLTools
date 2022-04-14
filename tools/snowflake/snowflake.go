// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 13:15
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : snowflake.go

package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

func Init() (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", "2022-04-01")
	if err != nil {
		return
	}

	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(1)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
