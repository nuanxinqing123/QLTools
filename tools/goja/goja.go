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
	"github.com/beego/beego/v2/adapter/httplib"
	"github.com/dop251/goja"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

// request HTTP请求方法
func request(wt interface{}, handles ...func(error, map[string]interface{}, interface{}) interface{}) interface{} {
	var method = "get"
	var url = ""
	var req *httplib.BeegoHTTPRequest
	var headers map[string]interface{}
	var formData map[string]interface{}
	var isJson bool
	var isJsonBody bool
	var body string
	var location bool
	var useproxy bool
	var timeout time.Duration = 0
	switch wt.(type) {
	case string:
		url = wt.(string)
	default:
		props := wt.(map[string]interface{})
		for i := range props {
			switch strings.ToLower(i) {
			case "timeout":
				timeout = time.Duration(Int64(props[i]) * 1000 * 1000)
			case "headers":
				headers = props[i].(map[string]interface{})
			case "method":
				method = strings.ToLower(props[i].(string))
			case "url":
				url = props[i].(string)
			case "json":
				isJson = props[i].(bool)
			case "datatype":
				switch props[i].(type) {
				case string:
					switch strings.ToLower(props[i].(string)) {
					case "json":
						isJson = true
					case "location":
						location = true
					}
				}
			case "body":
				if v, ok := props[i].(string); !ok {
					d, _ := json.Marshal(props[i])
					body = string(d)
					isJsonBody = true
				} else {
					body = v
				}
			case "formdata":
				formData = props[i].(map[string]interface{})
			case "useproxy":
				useproxy = props[i].(bool)
			}
		}
	}
	switch strings.ToLower(method) {
	case "post":
		req = httplib.Post(url)
	case "put":
		req = httplib.Put(url)
	case "delete":
		req = httplib.Delete(url)
	default:
		req = httplib.Get(url)
	}
	if timeout != 0 {
		req.SetTimeout(timeout, timeout)
	}
	if isJsonBody {
		req.Header("Content-Type", "application/json")
	}
	for i := range headers {
		req.Header(i, fmt.Sprint(headers[i]))
	}
	for i := range formData {
		req.Param(i, fmt.Sprint(formData[i]))
	}
	if body != "" {
		req.Body(body)
	}
	if location {
		req.SetCheckRedirect(func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		})
		rsp, err := req.Response()
		if err == nil && (rsp.StatusCode == 301 || rsp.StatusCode == 302) {
			return rsp.Header.Get("Location")
		} else
		//非重定向,允许用户自定义判断
		if len(handles) == 0 {
			return err
		}
	}
	if useproxy && Transport != nil {
		req.SetTransport(Transport)
	}
	rsp, err := req.Response()
	rspObj := map[string]interface{}{}
	var bd interface{}
	if err == nil {
		rspObj["status"] = rsp.StatusCode
		rspObj["statusCode"] = rsp.StatusCode
		data, _ := ioutil.ReadAll(rsp.Body)
		if isJson {
			zap.L().Debug("返回数据类型：JSON")
			var v interface{}
			json.Unmarshal(data, &v)
			bd = v
		} else {
			zap.L().Debug("返回数据类型：Not Is JSON")
			bd = string(data)
		}
		rspObj["body"] = bd
		h := make(map[string][]string)
		for k := range rsp.Header {
			h[k] = rsp.Header[k]
		}
		rspObj["headers"] = h
	}
	if len(handles) > 0 {
		return handles[0](err, rspObj, bd)
	} else {
		return bd
	}
}

// Request HTTP请求方法
func Request() interface{} {
	return request
}

// console 方法
var console = map[string]func(...interface{}){
	"info": func(v ...interface{}) {
		if len(v) == 0 {
			return
		}
		if len(v) == 1 {
			msg := fmt.Sprintf("Info: %s", v[0])
			fmt.Println(msg)
			return
		}
		msg := fmt.Sprintf("Info: %s", v)
		fmt.Println(msg)
	},
	"debug": func(v ...interface{}) {
		if len(v) == 0 {
			return
		}
		if len(v) == 1 {
			msg := fmt.Sprintf("Debug: %s", v[0])
			fmt.Println(msg)
			return
		}
		msg := fmt.Sprintf("Debug: %s", v)
		fmt.Println(msg)
	},
	"warn": func(v ...interface{}) {
		if len(v) == 0 {
			return
		}
		if len(v) == 1 {
			msg := fmt.Sprintf("Warn: %s", v[0])
			fmt.Println(msg)
			return
		}
		msg := fmt.Sprintf("Warn: %s", v)
		fmt.Println(msg)
	},
	"error": func(v ...interface{}) {
		if len(v) == 0 {
			return
		}
		if len(v) == 1 {
			msg := fmt.Sprintf("Error: %s", v[0])
			fmt.Println(msg)
			return
		}
		msg := fmt.Sprintf("Error: %s", v)
		fmt.Println(msg)
	},
	"log": func(v ...interface{}) {
		if len(v) == 0 {
			return
		}
		if len(v) == 1 {
			msg := fmt.Sprintf("Info: %s", v[0])
			fmt.Println(msg)
			return
		}
		msg := fmt.Sprintf("Info: %s", v)
		fmt.Println(msg)
	},
}

// Int64 转换Int64
var Int64 = func(s interface{}) int64 {
	i, _ := strconv.Atoi(fmt.Sprint(s))
	return int64(i)
}
