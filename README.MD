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
- （未完成）容器独享变量限额
- 自选服务器和变量组上传
- （未完成）支持容器相互迁移&复制
- JWT认证
- Limiter保障服务器访问频率和次数
- SQLite轻量数据库（无需繁杂配置）
- 支持优雅重启（处理完已接收的请求后再关闭服务器）

## 🍳技术栈
- 语言：Golang
- Web框架：Gin框架
- 配置文件：Viper库
- 日志：Zap库
- 数据库：GORM - SQLite
- API文档：暂不公布
- 文档密码：暂不公布


## 🗽待修复Bug
- 后台侧边导航栏状态异常
- 后台侧边导航栏手机端切换页面后出现灰色蒙版


### 静态资源构建
```shell
go-bindata -o=bindata/bindata.go -pkg=bindata ./assets/...
```