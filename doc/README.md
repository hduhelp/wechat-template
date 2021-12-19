# 项目开发说明 开发文档

## 环境初始化

[环境初始化与排错文档](./developEnvDebug.md)

## Commit 规范

### 标准格式
```
Action(modname/package): (add|update|..all action you done)
```

### example

```
update(wechatPong): add template message test 
```

## 接口标准

1. URL中字母全部小写
2. 如果有单词拼接，使用中划线‘-’，不使用下划线‘_’ (在搜索引擎中，把中划线当做空格处理，而下划线是被忽略的。使用中划线是对搜索引擎友好的写法)
3. 资源必须采用资源名词的复数形式
4. 层级尽量不超过三层
5. 参数不允许超过3个
6. 如果是内容资源的URL，不允许以参数方式显示

## 项目结构 

> WorkingDir . is repo root

### ./server

服务基础软件 所在目录 中间件

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

- database.go 数据库的重启和创建
- messageHandler.go 微信公众号获取的消息处理
  - 重要结构体:
    - RouterGroup
    - MsgEngine
    - Message
  - 主要别名:
    - HandlerFunc
    - HandlersChain
  - 主要方法:
    - MsgEngine.Group 
    - MsgEngine.Use
    - MsgEngine.MsgText
    - MsgEngine.EventClick
    - 更多请自行查看 IRoutes 接口
- server.go 进行服务的初始化
  - 重要结构体 
    - Server 包含了 微信引擎 gin 消息处理函数
  - 主要函数
    - 服务的初始化 启动 结束
- wechatBind.go
  - 重要结构体
  - 重要函数 进行 hduid 与 微信 id 的相互转换
    - ConvertWechatUnionIdToHDUId 
    - ConvertHDUIdToWechatUnionId


### ./modules

模块文件夹 使用文件夹 集合每一个子模块的 go 模块 package 表示

并且在 [main.go](../main.go) 启用你的模块

[示例模块](../module/wechatPong/init.go)

### ./server

框架文件夹

## 微信公众号配置 

[配置](./developEnvDebug.md)