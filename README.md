# wechat-template

杭电助手目前正在使用的微信公众号后端速冲框架，便于小型业务的敏捷开发，是一个多模块组合设计的微信公众号后端程序。

包装了基础功能,同时设计了一个~~良好~~的项目结构

## 快速开始

1. 项目配置：按照示例配置文件 [./application.example.yaml](./application.example.yaml) 新建配置文件 application.yaml 
2. 运行本框架 
3. 微信端配置：前往微信后端填写服务器地址(URL)为「你配置的地址 + `/serve`」，如：
    > http://xxx.xxxx.com/serve

如果一切顺利，此时你的公众号后端就上线啦。  

可通过向公众号发送 `ping` 判断是否走通，正确配置将收到回复。

## 快速编写子模块

1. 在 `/module` 下新建文件夹，内容可参考 [wechatPong](./module/wechatPong/init.go)
2. 开启子模块：在 [main.go](./main.go) 中以 `import _` 的方式引用模块 package，即可启用Module

完整可参考：[自定义组件方法与组件插拔逻辑](./doc/moduleRegister.md)


## 开发文档详细指南
[开发文档与规范](./doc/)

## 不了解go?

golang 极速入门

[点我看书](https://github.com/justjavac/free-programming-books-zh_CN#go)



## Module 配置

module参考[wechatPong](https://github.com/hduhelp/wechat-template/blob/main/module/wechatPong/init.go)

编写自己的Module后在 [main.go](https://github.com/hduhelp/wechat-template/blob/main/main.go) 中启用Module

[自定义组件方法与组件插拔逻辑](https://github.com/hduhelp/wechat-template/tree/main/doc/moduleRegister.md)

