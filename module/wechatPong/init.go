package wechatPong

import (
	"github.com/getsentry/sentry-go"
	"github.com/hduhelp/wechat-template/server"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"sync"
)

func init() {
	instance = &wechatPong{}
	server.RegisterModule(instance)
}

var instance *wechatPong

type wechatPong struct {
}

func (m *wechatPong) GetModuleInfo() server.ModuleInfo {
	return server.ModuleInfo{
		// module 的 id 全部用 NewModule 函数生成
		// namespace 为开发者，moduleName 为模块名（一般用包名） 方便锤开发者以及定位出错的包
		ID:       server.NewModuleID("atom", "wechatPong"),
		Instance: instance,
	}
}

func (m *wechatPong) Init() {
	// 初始化过程
	// 在此处可以进行 Module 的初始化配置
	// 如配置读取
}

func (m *wechatPong) PostInit() {
	// 第二次初始化
	// 再次过程中可以进行跨Module的动作
	// 如通用数据库等等
}

func (m *wechatPong) Serve(s *server.Server) {
	// 注册服务函数部分
	// index 是匹配的优先级，index越大，优先级越高，优先测试该条匹配规则
	// key 就是正则匹配的规则 当匹配上之后就该条路由
	// 可以在 https://regex101.com/ 测试你的正则规则
	s.MsgEngine.Group("^ping$", func(msg *server.Message) {
		//// 下面是调用模组发送模板消息
		//templateMessageModule, _ := server.GetModule(server.NewModuleID("atom", "templateMessage"))
		//templateMessageSender := templateMessageModule.Instance.(*templateMessage.Module)
		//templateMessageSender.PushMessage(&templateMessage.TemplateMessage{
		//	Message: &message.TemplateMessage{
		//		ToUser:     msg.GetOpenID(),
		//		TemplateID: "请填写你可用的模板的id",
		//		Data: map[string]*message.TemplateDataItem{
		//			"keyword1": {
		//				Value: "ping-pong模块",
		//			},
		//			"keyword2": {
		//				Value: "模板消息模组测试",
		//			},
		//		},
		//	},
		//	Resend: true,
		//})

		msg.Reply = &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("pong\n" +
			"Current version: " + server.Version)}
	}). //下面就是注册路由的部分，可以像下面一样同时注册文字消息和按钮点击事件。还有更多使用方式请查看框架中实现的接口。
		MsgText("", 1).EventClick("")

	// 由于group已经定义了一个中间件实现pong功能 所以后面注册具体方法的时候并不需要带上查询函数
	// 同时由于路由是由baseKey+key拼接而成，所以也不需要单独设置key了
	// 这样还可以快速实现一个查询函数多个查询关键字，如：
	s.MsgEngine.Group("", func(msg *server.Message) {
		msg.Reply = &message.Reply{
			MsgType: message.MsgTypeText,
			MsgData: message.NewText("ping"),
		}
	}).MsgText("^pong$", 1).MsgText("^poooong$", 1)
	// 试试向你的公众号发送pong或poooong看看效果吧
}

func (m *wechatPong) Start(server *server.Server) {
	// 请在可能出错的地方使用 sentry 接住错误 越早defer越好
	defer sentry.Recover()

	// 此函数会新开携程进行调用
	// ```go
	// 		go exampleModule.Start()
	// ```

	// 可以利用此部分进行后台操作
	// 如http服务器等等
}

func (m *wechatPong) Stop(server *server.Server, wg *sync.WaitGroup) {
	// 别忘了解锁
	defer wg.Done()
	// 结束部分
	// 一般调用此函数时，程序接收到 os.Interrupt 信号
	// 即将退出
	// 在此处应该释放相应的资源或者对状态进行保存
}
