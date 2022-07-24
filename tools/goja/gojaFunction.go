// -*- coding: utf-8 -*-
// @Time    : 2022/6/8 21:28
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : gojaFunction.go

package goja

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/adapter/httplib"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

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

// refind 正则表达式
func refind(re, data string) interface{} {
	reg := regexp.MustCompile(re)
	return reg.FindAllStringSubmatch(data, -1)
}

// ReFind 正则表达式
func ReFind() interface{} {
	return refind
}

// Replace 字符串替换
func Replace(s, old, new string, count int) interface{} {
	c := 1
	if count != 0 {
		c = count
	}
	return strings.Replace(s, old, new, c)
}
