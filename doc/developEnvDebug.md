# 环境建立与排错

## 第三方依赖

- [wechat golang api 库](https://github.com/silenceper/wechat)

其中 以上项目 README.md 呈现了相应的环境配置方法 甚至简单的使用用例

PS: **需要注意** 如果遇到 微信 API 并未实现的 可参照对应的 API 规则进行请求

## 项目初始化

请注意在 go 1.17 或者以上版本的情况下进行处理

在项目的环境情况下 (存在 go.mod) 进行 `go mod tidy` 操作

会自动从远程拉取项目 若遇到 404、 timeout 等错误 可能是您的代理爆炸了

您需要检查 代理配置

- [wechat API official Document](https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html)

## 配置文件

配置文件准备：复制 [配置文件 样本](../application.example.yaml) 并重命名为 `application.yaml`。

### 项目配置
目前项目配置仅设置了 `gin` 端口： `httpEngine.port`，默认 9955 。

### 微信公众号测试号配置

#### 测试号申请与配置

前往 [微信官方](https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login) 申请测试公众号，  
在网页内填写相应信息，并修改对应的配置，详见下图：  

![配置微信app id 和 secret](./img/wx_app_url_token_conf.png)  

下为 `$Project_Root/application.yaml` 的 `wechat` 配置部分  
```yaml
wechat:
  appID: ...见图描述
  appSecret: ...见图描述
  token: ...见图描述
  encodingAESKey: ... 留空即可
```

Just like [文档](https://silenceper.com/wechat/officialaccount/start.html) 中所提到的
在部署 测试性质服务器的时候 可以使用 ngrok 类的反向代理 生成一个 反代 然后使用自己的本地启动的服务来进行公众号的调试

如果您有自己的服务器域名 与 SSL 证书可以完全使用正常的服务器进行测试 这可能相对而言较为麻烦 当然更加贴近实际的生产环境

#### 这里是小小的 ngrok 使用演示 

Q：为什么要用 ngrok ？

A：用户操作微信客户端，请求会先发送至微信服务器，再由微信服务器请求本项目。所以本项目需要提供一个公网可访问的服务，用于接收微信服务器的请求。而我们本地测试环境无公网ip，所以需要使用ngrok进行内网穿透。

Windows 下可以快速使用 [ngrok Win amd64.exe](https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-windows-amd64.zip) (无需注册帐号...ngrok 的官方政策是 只有 https 等链接才需要账户) 

操作如下 所示

- setup the ngrok ![step1](img/ngrok_conf.png)
  - command line: .\ngrok.exe http 9955 # 请注意 这里 9955 是根据你的 application.yaml 中 httpEngine port 规定的 example 中是 9955
- config the url like ![step2](img/url_and_token.png)
  - **请注意** 默认的 api 位置在 /serve 下 (根据 [代码](https://github.com/hduhelp/wechat-template/blob/58819b1250368716726054648fb1a9235811c194/server/server.go#L92) 110 行 StartService() 函数的定义)

然后进行编译和启动本项目就可以在微信官方平台下 快速通过 api 域名的验证

都完成了之后 你就得到了 一个自己的微信公众号后端







