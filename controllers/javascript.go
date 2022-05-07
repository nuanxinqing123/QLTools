// -*- coding: utf-8 -*-
// @Time    : 2022/5/6 16:04
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : javascript.go

package controllers

import (
	"QLPanelTools/model"
	res "QLPanelTools/tools/response"
	val "QLPanelTools/tools/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// JavascriptUpload 上传插件
func JavascriptUpload(c *gin.Context) {
	// 获取上传文件
	file, err := c.FormFile("file")
	if err != nil {
		res.ResError(c, res.CodeServerBusy)
		return
	}

	// 获取插件目录绝对路径
	ExecPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		zap.L().Error(err.Error())
		res.ResError(c, res.CodeServerBusy)
		return
	}

	// 保存文件
	FilePath := ExecPath + "/plugin/" + file.Filename
	err = c.SaveUploadedFile(file, FilePath)
	if err != nil {
		zap.L().Error(err.Error())
		res.ResError(c, res.CodeServerBusy)
		return
	}

	res.ResSuccess(c, res.CodeSuccess)
}

// JavascriptDelete 删除插件
func JavascriptDelete(c *gin.Context) {
	// 获取参数
	p := new(model.DeletePlugin)
	if err := c.ShouldBindJSON(&p); err != nil {
		// 参数校验
		zap.L().Error("SignInHandle with invalid param", zap.Error(err))

		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			res.ResError(c, res.CodeInvalidParam)
			return
		}

		// 翻译错误
		res.ResErrorWithMsg(c, res.CodeInvalidParam, val.RemoveTopStruct(errs.Translate(val.Trans)))
		return
	}

	// 获取插件目录绝对路径
	ExecPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		zap.L().Error(err.Error())
		res.ResError(c, res.CodeServerBusy)
		return
	}

	// 删除文件
	FilePath := ExecPath + "/plugin/" + p.FileName
	err = os.Remove(FilePath)
	if err != nil {
		// 删除失败
		zap.L().Error(err.Error())
		res.ResError(c, res.CodeRemoveFail)
		return
	}

	res.ResSuccess(c, res.CodeSuccess)
}

// JavascriptReadall 获取插件目录下所有插件
func JavascriptReadall(c *gin.Context) {
	// 获取插件目录绝对路径
	ExecPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		zap.L().Error(err.Error())
		res.ResError(c, res.CodeServerBusy)
		return
	}

	// 删除文件
	PluginPath := ExecPath + "/plugin/"

	// 读取目录
	var fl []model.FileInfo
	var fi model.FileInfo
	files, _ := ioutil.ReadDir(PluginPath)
	for _, f := range files {
		// 跳过不是JS的文件
		if !strings.Contains(f.Name(), ".js") {
			continue
		}

		zap.L().Debug(f.Name())

		// 读取插件名称
		fd, err := os.Open(PluginPath + f.Name())
		if err != nil {
			zap.L().Error(f.Name() + "：打开文件失败")
		}
		defer fd.Close()
		v, _ := ioutil.ReadAll(fd)
		data := string(v)
		FileIDName := ""
		if regs := regexp.MustCompile(`\[name:(.+)]`).FindStringSubmatch(data); len(regs) != 0 {
			FileIDName = strings.Trim(regs[1], " ")
		}

		fi.FileName = f.Name()
		fi.FileIDName = FileIDName
		fl = append(fl, fi)
	}

	res.ResSuccess(c, fl)
}
