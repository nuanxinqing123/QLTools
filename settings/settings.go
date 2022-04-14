// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 12:55
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : settings.go

package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type App struct {
	// 运行模式
	Mode string `mapstructure:"mode"`
	// 运行端口
	Port int `mapstructure:"port"`
}

// Conf 保存配置信息
var Conf = new(App)

func Init() (err error) {
	// 配置文件名称
	viper.SetConfigName("config")
	// 配置文件扩展名
	viper.SetConfigType("yaml")
	// 配置文件所在路径
	viper.AddConfigPath("./config")
	// 查找并读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		// 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}

	// 配置信息绑定到结构体变量
	err = viper.Unmarshal(Conf)
	if err != nil {
		fmt.Printf("viper.Unmarshal() failed, err:%v\n", err)
	}

	// 热加载配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(c fsnotify.Event) {
		fmt.Println("检测到配置文件有变动,已实时加载")
	})
	return
}
