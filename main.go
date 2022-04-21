// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 12:46
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : main.go

package main

import (
	"QLPanelTools/logger"
	"QLPanelTools/routes"
	"QLPanelTools/settings"
	"QLPanelTools/sqlite"
	"QLPanelTools/tools/snowflake"
	"QLPanelTools/tools/validator"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	// 判断注册默认配置文件
	bol := IFConfig("config/config.yaml")
	if bol != true {
		fmt.Println("自动生成配置文件失败, 请按照仓库内容手动在config/config.yaml下创建配置文件")
		return
	}

	// 加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("viper init failed, err:%v\n", err)
		return
	}

	// 初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("Logger init failed, err:%v\n", err)
		return
	}
	defer func(l *zap.Logger) {
		_ = l.Sync()
	}(zap.L())
	zap.L().Debug("Logger success init ...")

	// 初始化数据库
	sqlite.Init()
	zap.L().Debug("SQLite success init ...")

	// 初始化翻译器
	if err := validator.InitTrans("zh"); err != nil {
		fmt.Printf("validator init failed, err:%v\n", err)
		return
	}
	zap.L().Debug("Validator success init ...")

	// 注册雪花ID算法
	if err := snowflake.Init(); err != nil {
		fmt.Printf("snowflake init failed, err:%v\n", err)
		return
	}
	zap.L().Debug("Snowflake success init ...")

	// 运行模式
	if viper.GetString("app.mode") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 注册路由
	r := routes.Setup()

	// 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	fmt.Println("服务监听端口:" + strconv.Itoa(viper.GetInt("app.port")))
	zap.L().Info("服务监听端口:" + strconv.Itoa(viper.GetInt("app.port")))

	// 优雅关机
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listten: %s\n", err)
		}
	}()

	// 等待终端信号来优雅关闭服务器，为关闭服务器设置5秒超时
	quit := make(chan os.Signal, 1) // 创建一个接受信号的通道

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞此处，当接受到上述两种信号时，才继续往下执行
	zap.L().Info("Service ready to shut down")

	// 创建五秒超时的Context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 五秒内优雅关闭服务（将未处理完成的请求处理完再关闭服务），超过五秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Service timed out has been shut down：", zap.Error(err))
	}

	zap.L().Info("Service has been shut down")
}

type Config struct {
	App App `yaml:"app"`
}

type App struct {
	Mode string `yaml:"mode"`
	Port int    `yaml:"port"`
}

// IFConfig 判断并自动生成启动配置文件
func IFConfig(src string) bool {
	// 检测是否存在配置文件
	file, err := os.Stat(src)
	if err != nil {
		_, err = os.Create(src)
		if err != nil {
			fmt.Printf("Create Config File Error: %s", err)
			return false
		}

		// 需要生成默认
		cfg := &Config{
			App: App{
				Mode: "",
				Port: 15000,
			},
		}
		data, err := yaml.Marshal(cfg)
		if err != nil {
			fmt.Printf("Marshal Config Error: %s", err)
			return false
		}

		// 写入默认配置
		err = ioutil.WriteFile(src, data, 0777)
		if err != nil {
			fmt.Printf("WriteFile Config Error: %s", err)
			return false
		}

		return true
	}
	if file.IsDir() {
		return false
	} else {
		return true
	}
}
