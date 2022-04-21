<!--suppress HtmlDeprecatedAttribute -->
<p align="center">
  <a href="https://github.com/whyour/qinglong">
    <img width="150" src="https://z3.ax1x.com/2021/11/18/I7MpAe.png" alt="Img">
  </a>
</p>

<h1 align="center">青龙Tools</h1>

## 🍭功能介绍
- 自定义变量名称
- 自定义变量的数量
- 支持多容器
- 容器独享变量限额
- 自选服务器和变量组上传
- （未完成）支持容器相互迁移&复制
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
1、直接启动
```shell
# 创建目录,自己夺定合适上级目录
mkdir qltools && cd qltools
# 赋予权限
chmod 755 QLPanelTools
# 上传程序二进制文件后启动
./QLPanelTools
```
2、Docker启动
```shell
目前为内/公测阶段, 不放出Docker Images
```

## 🗽待修复Bug
- 暂无


### 静态资源构建
```shell
go-bindata -o=bindata/bindata.go -pkg=bindata ./assets/...
```