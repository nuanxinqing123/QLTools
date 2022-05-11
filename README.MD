<!--suppress HtmlDeprecatedAttribute -->
<p align="center">
  <a href="https://github.com/whyour/qinglong">
    <img width="150" src="https://z3.ax1x.com/2021/11/18/I7MpAe.png" alt="Img">
  </a>
</p>

<h1 align="center">青龙Tools</h1>
<p><a href="https://github.com/nuanxinqing123/QLTools#%E6%BC%94%E7%A4%BA%E5%9B%BE">查看演示图</a></p>

## 🍭功能介绍
- 自定义变量名称 & 数量
- 支持多容器上传管控
- 容器独享变量限额
- 自选服务器和变量组上传
- 支持变量黑名单 & IP提交次数限制
- 支持容器相互迁移 & 复制 & 备份 & 恢复
- JWT认证
- Limiter保障服务器访问频率
- SQLite轻量数据库（无需繁杂配置）
- 支持优雅关机（处理完已接收的请求后再关闭服务器）

## 🍳技术栈
- 语言：Golang
- Web框架：Gin框架
- 配置文件：Viper库
- 日志：Zap库
- 数据库：GORM - SQLite
- API文档：https://console-docs.apipost.cn/preview/0fdb4c815ed24ab2/e2057f0e0b8dc545
- 文档密码：609889

## 🧸配置文件
```yaml
app:
  # 运行模式：生产环境留空(开发模式：debug)
  mode: ""
  # 运行端口
  port: 15000
```

## 🔍安装教程
1、Supervisord & PM2 & nohup 启动守护
```shell
# 创建目录并进入
mkdir QLTools && cd QLTools

# 给予权限
chmod 755 程序名称

# 启动程序排查启动错误
./程序名称

# 程序无误后点击下方教程查看程序后台守护教程
```

教程地址：[青龙Tools的后台进程守护教程（附反代域名）](https://6b7.org/460.html)

2、Docker启动
```shell
# 创建QLTools目录并进入
mkdir qltools && cd qltools

# Docker版本提供架构：amd64、arm64、arm-7
docker run -itd --name QLTools -v $PWD/config:/QLTools/config -v $PWD/plugin:/QLTools/plugin -p 15000:15000 nuanxinqing123/qltools:latest

# 更新步骤：因为当前版本不支持热更新，所以需要暂停并删除容器，重新拉取新版本镜像后在qltools目录里面重新启动
```

## 🎯开发计划

开发计划 & 进度：[点击查看](https://web.banlikanban.com/kanban/626f9b4c6ade1220282ac551)

招募一位热心愿意一起开发的前端小伙伴，嘿嘿。我前端太差了

TG：[https://t.me/Nuanxinqing](https://t.me/Nuanxinqing)

Email：nuanxinqing@gmail.com

## 🧩当前版本日志
- 修复 前端：Select选项切换数据错误
- 修复 前端：Select显示数据错误
- 修复 前端：切换服务器数据并未切换错误

## 📔自行构建步骤
```shell
# 项目开发环境：Golang 1.18
# Clone 项目并进入项目目录
git clone https://github.com/nuanxinqing123/QLTools.git && cd QLTools

# 更新项目依赖
go mod tidy

# 构建前端静态文件（需要安装go-bindata）
go-bindata -o=bindata/bindata.go -pkg=bindata ./assets/...

# 打包
go build
```

## 📷演示图

![QQ截图20220511154438.png](https://pic.6b7.xyz/2022/05/11/25a5e41170f5f.png)

![QQ截图20220511154454.png](https://pic.6b7.xyz/2022/05/11/3f13f15d25b46.png)

![QQ截图20220511154818.png](https://pic.6b7.xyz/2022/05/11/e41ea41542307.png)

![QQ截图20220511154933.png](https://pic.6b7.xyz/2022/05/11/40f36ef85f79d.png)

![QQ截图20220511154947.png](https://pic.6b7.xyz/2022/05/11/347a5fd9b12f2.png)

![QQ截图20220511155004.png](https://pic.6b7.xyz/2022/05/11/3c3c339fa3b82.png)

![QQ截图20220511155021.png](https://pic.6b7.xyz/2022/05/11/4fe5dab516d93.png)



