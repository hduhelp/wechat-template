# 项目开发说明/开发文档

## 环境初始化

[环境初始化与排错文档](./developEnvDebug.md)


## 建议与规范
[建议与规范](./spec.md)

## 项目结构 

> WorkingDir . is repo root

### 框架核心层 `./server`

实现了模块化的统一注册管理逻辑与微信消息链的封装。如无必要，不应更改此模块。

- server.go 组织框架核心
- module.go 文件规定了模块的 必要函数
  
  下图 按照执行顺序排列
  >
  > init 模块初始化
  > 
  > PostInit 二次初始化
  >
  > Serve main 区域
  > 
  > Start server 的开始区 可以 defer sentry.Recover() 用以捕捉错误
  >
  > Stop() 结束时操作 一般为 defer wg.Done 
- middleware.go 中间件
- moduleId.go 模块id
- moduleInfo.go 模块信息
- msgContext.go 消息上下文
- msgEngine.go 微信消息引擎


### 模块层 `./modules`

模块文件夹，此处编写业务逻辑。每一个文件夹即为一个 `go package`，也即一个模块。

[模块详细介绍](./moduleRegister.md)
