# wechat-template

杭电助手目前正在使用的微信公众号后端速冲框架，便于小型业务的敏捷开发，是一个多模块组合设计的微信公众号后端程序。

包装了基础功能,同时设计了一个~~良好~~的项目结构

## 快速开始

按照示例配置文件 application.example.yaml 新建一个配置文件 application.yaml

然后运行本框架

前往微信后端填写服务器地址(URL)为你配置的地址+/serve，如：
> http://xxx.xxxx.com/serve

如果一切顺利，此时你的公众号后端就上线啦。

## [开发文档](./doc/)

## 不了解go?

golang 极速入门

[点我看书](https://github.com/justjavac/free-programming-books-zh_CN#go)



## Module 配置

module参考[wechatPong](https://github.com/hduhelp/wechat-template/blob/main/module/wechatPong/init.go)

编写自己的Module后在 [main.go](https://github.com/hduhelp/wechat-template/blob/main/main.go) 中启用Module

[自定义组件方法与组件插拔逻辑](https://github.com/hduhelp/wechat-template/tree/main/doc/moduleRegister.md)

