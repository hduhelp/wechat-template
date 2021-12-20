# wechat-template

杭电助手目前正在使用的微信公众号后端速冲框架，便于小型业务的敏捷开发，是一个多模块组合设计的微信公众号后端程序。

包装了基础功能,同时设计了一个~~良好~~的项目结构

## 框架优点
1. 模块化的设计，使得模块间几乎无干扰，非常适合多人协作速冲。传统项目新人需要看懂整个项目才敢下手，此框架下，新人只需要关注自己的模块。  
2. 普通开发者无需关注 server 层实现逻辑，依葫芦画瓢实现相应接口即可。

注意事项：尽可能减少模块间引用，防止循环引用 `import circle :(`。

## 快速开始

1. 项目配置：按照示例配置文件 [application.example.yaml](./application.example.yaml) 新建配置文件 `application.yaml` 。（[完整环境配置文档](./doc/developEnvDebug.md))
2. 运行本框架 
3. 微信端配置：前往微信公众号网页填写。其中，服务器地址(URL)为「你配置的地址 + `/serve`」，如：
    > http://xxx.xxxx.com/serve

如果一切顺利，此时你的公众号后端就上线啦。  

可通过向公众号发送 `ping` 判断是否走通，正确配置将收到回复。

## 快速编写子模块

1. 在 `/module` 下新建文件夹，内容可参考 [wechatPong](./module/wechatPong/init.go)
2. 开启子模块：在 [main.go](./main.go) 中以 `import _` 的方式引用模块 package，即可启用Module

完整可参考：[自定义组件方法与组件插拔逻辑](./doc/moduleRegister.md)


## 开发文档详细指南
[开发文档详细指南](./doc/README.md)

## 不了解go?

1. [HDUHELP Official Backend Document](https://github.com/hduhelp/backend_guide/)
2. [点我看书](https://github.com/justjavac/free-programming-books-zh_CN#go)

