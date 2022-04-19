// -*- coding: utf-8 -*-
// @Time    : 2022/4/2 13:18
// @Author  : Nuanxinqing
// @Email   : nuanxinqing@gmail.com
// @File    : routes.go

package routes

import (
	"QLPanelTools/bindata"
	"QLPanelTools/controllers"
	"QLPanelTools/middleware"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
	"time"
)

func Setup() *gin.Engine {
	// 创建服务
	r := gin.Default()

	// 配置中间件
	{
		//r.Use(logger.GinLogger(), logger.GinRecovery(true))
		// 限流熔断
		r.Use(middleware.RateLimitMiddleware(time.Minute, 500, 500)) // 每分钟限制500次请求, 超出熔断
	}

	// 前端静态文件
	{
		// 加载模板文件
		t, err := loadTemplate()
		if err != nil {
			panic(err)
		}
		r.SetHTMLTemplate(t)

		// 加载静态文件
		fs := assetfs.AssetFS{
			Asset:     bindata.Asset,
			AssetDir:  bindata.AssetDir,
			AssetInfo: nil,
			Prefix:    "assets",
		}
		r.StaticFS("/static", &fs)

		r.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", nil)
		})
	}

	// 路由组
	{
		// 开放权限组
		open := r.Group("v1/api")
		{
			// 账户注册
			open.POST("signup", controllers.SignUpHandle)
			// 账户登录
			open.POST("signin", middleware.RateLimitMiddleware(time.Minute, 10, 10), controllers.SignInHandle) // 每分钟限制10次请求, 超出熔断
			// 检查Token是否有效
			open.POST("check/token", controllers.CheckToken)

			// 可用服务
			open.GET("index/data", controllers.IndexData)
			// 上传变量
			open.POST("env/add", controllers.EnvADD)
		}

		// 管理员权限组
		ad := r.Group("v2/api")
		ad.Use(middleware.UserAuth())
		{
			// 测试
			ad.GET("123", controllers.AdminTest)
			// 面板连接测试
			ad.POST("panel/connect", controllers.GetPanelToken)

			// 管理员：获取前十次登录记录
			ad.GET("admin/ip/info", controllers.GetIPInfo)
			// 管理员：密码修改
			ad.POST("admin/rep-wd", controllers.ReAdminPwd)
			// 管理员: 获取管理员信息
			ad.GET("admin/info", controllers.GetAdminInfo)

			// 变量名：新增
			ad.POST("env/name/add", controllers.EnvNameAdd)
			// 变量名：修改
			ad.PUT("env/name/update", controllers.EnvNameUp)
			// 变量名：删除
			ad.DELETE("env/name/del", controllers.EnvNameDel)
			// 变量名：All
			ad.GET("env/name/all", controllers.GetAllEnvData)

			// 面板：新增
			ad.POST("env/panel/add", controllers.PanelAdd)
			// 面板：修改
			ad.PUT("env/panel/update", controllers.PanelUp)
			// 面板：删除
			ad.DELETE("env/panel/del", controllers.PanelDel)
			// 面板：All
			ad.GET("env/panel/all", controllers.GetAllPanelData)
			// 面板：绑定变量
			ad.PUT("env/panel/binding/update", controllers.UpdatePanelEnvData)
		}
	}

	return r
}

//加载模板文件
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for _, name := range bindata.AssetNames() {
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		asset, err := bindata.Asset(name)
		if err != nil {
			continue
		}
		name := strings.Replace(name, "assets/", "", 1)
		t, err = t.New(name).Parse(string(asset))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
