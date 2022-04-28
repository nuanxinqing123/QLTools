<!--suppress HtmlDeprecatedAttribute -->
<p align="center">
  <a href="https://github.com/whyour/qinglong">
    <img width="150" src="https://z3.ax1x.com/2021/11/18/I7MpAe.png" alt="Img">
  </a>
</p>

<h1 align="center">青龙Tools</h1>
<p><a href="https://github.com/nuanxinqing123/QLTools#%E6%BC%94%E7%A4%BA%E5%9B%BE">查看演示图</a></p>

## 🍭功能介绍
- 自定义变量名称
- 自定义变量的数量
- 支持多容器
- 容器独享变量限额
- 自选服务器和变量组上传
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

## 🎯开发计划
- 消息推送（推送敏感操作记录）

## 🧩当前版本日志
- 修复 PC端首页Footer部分不会自动拉伸
- 新增 容器相互迁移&复制
- 新增 变量重复上传检测
- 新增 变量新建&合并模式
- 新增 版本更新检查
- 更新 上传变量的提示改为弹窗提示

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

![1.png](https://pic.6b7.xyz/2022/04/22/b100b8eaffafc.png)

![2.png](https://pic.6b7.xyz/2022/04/22/ba035ba2c99e3.png)

![3.png](https://pic.6b7.xyz/2022/04/22/411faac7d8f55.png)

![4.png](https://pic.6b7.xyz/2022/04/22/e0e76237e9aa7.png)

![5.png](https://pic.6b7.xyz/2022/04/22/4473e0a153112.png)



