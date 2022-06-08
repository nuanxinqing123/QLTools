// -*- coding: utf-8 -*-
// @Time    : 2022/5/5 21:37
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : goja.go

package goja

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dop251/goja"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var Transport *http.Transport

/*
1、创建Request方法
2、创建goja方法
3、向goja传入JS路径和变量
4、判断传入数据是不是JS文件或是不是一个文件夹
5、执行完成必须返回数据：bool(判断是否允许提交), int(1、新建 - 2、合并 - 3、更新), string(处理完成后的CK), error(函数执行中的错误)
*/

type jsonData struct {
	Bool bool   `json:"bool"`
	Env  string `json:"env"`
}

// RunJS 执行javascript代码
func RunJS(filename, env string) (bool, string, error) {
	// 获取运行的绝对路径
	ExecPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		zap.L().Error(err.Error())
		return false, "", err
	}

	if !strings.Contains(filename, ".js") {
		// 不是JS文件
		return false, "", errors.New("传入值不是JS文件名")
	}

	// JS文件完整路径
	JSFilePath := ExecPath + "/plugin/" + filename
	// 读取文件内容
	data, err := os.ReadFile(JSFilePath)
	if err != nil {
		zap.L().Error(err.Error())
		return false, "", err
	}
	template := string(data)

	// 创建JS虚拟机
	vm := goja.New()
	// 注册JS方法
	vm.Set("Request", Request)
	vm.Set("request", request)
	vm.Set("console", console)
	vm.Set("refind", refind)
	vm.Set("ReFind", ReFind)
	_, err = vm.RunString(template)
	if err != nil {
		// JS代码有问题
		zap.L().Error(err.Error())
		return false, "", err
	}
	var mainJs func(string) interface{}
	err = vm.ExportTo(vm.Get("main"), &mainJs)
	if err != nil {
		// JS函数映射到 Go函数失败
		zap.L().Error(err.Error())
		return false, "", err
	}

	var j jsonData
	jd := mainJs(env)
	marshal, err := json.Marshal(jd)
	if err != nil {
		return false, "", err
	}
	json.Unmarshal(marshal, &j)

	if j.Bool {
		zap.L().Debug("true")
	} else {
		zap.L().Debug("false")
	}
	zap.L().Debug(j.Env)

	return j.Bool, j.Env, nil
}

// Int64 转换Int64
var Int64 = func(s interface{}) int64 {
	i, _ := strconv.Atoi(fmt.Sprint(s))
	return int64(i)
}
