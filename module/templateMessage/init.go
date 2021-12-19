// 集中管理发送模板消息的模组
// 消息发送相关的配置请在 method.go 中修改

package templateMessage

import (
	"github.com/hduhelp/wechat-template/server"
	"github.com/hduhelp/wechat-template/utils"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/sirupsen/logrus"
	"sync"
)

func init() {
	instance = &Module{}
	server.RegisterModule(instance)
	logger = utils.GetModuleLogger(instance.GetModuleInfo().String())
}

var logger *logrus.Entry

var instance *Module

type Module struct {
	template               *message.Template
	MessageQueue           chan *TemplateMessage
	MessageSenderWaitGroup sync.WaitGroup
	senderLimit            chan struct{} // 最大协程限制
}

func (m *Module) GetModuleInfo() server.ModuleInfo {
	return server.ModuleInfo{
		ID:       server.NewModuleID("atom", "templateMessage"),
		Instance: instance,
	}
}

func (m *Module) Init() {
	logger.Info("Init template message sender...")

	m.MessageQueue = make(chan *TemplateMessage, maxSenderNumber)
	m.senderLimit = make(chan struct{}, maxSenderNumber)
}

func (m *Module) PostInit() {

}

func (m *Module) Serve(s *server.Server) {
	m.template = message.NewTemplate(s.WechatEngine.GetContext())
	go m.registerMessageSender(m.MessageQueue)
}

func (m *Module) Start(s *server.Server) {
	// example:
	//m.PushMessage(&TemplateMessage{
	//	Message: &message.TemplateMessage{
	//		ToUser:     "unionId",
	//		TemplateID: "_HtuD7TrFKxquwizJwICXv4sWg5AeZBvaHBIRvYKeKk",
	//		URL:        "",
	//		Color:      "",
	//		Data: map[string]*message.TemplateDataItem{
	//			"keyword1": &message.TemplateDataItem{
	//				Value: "测试消息",
	//				Color: "",
	//			},
	//			"keyword2": &message.TemplateDataItem{
	//				Value: "中午问候",
	//				Color: "",
	//			},
	//		},
	//		MiniProgram: struct {
	//			AppID    string `json:"appid"`
	//			PagePath string `json:"pagepath"`
	//		}{},
	//	},
	//	Resend:       true,
	//	RetriedTime:  0,
	//	MaxRetryTime: 0,
	//})
}

func (m *Module) Stop(s *server.Server, wg *sync.WaitGroup) {
	close(m.senderLimit)
	close(m.MessageQueue) // 关闭此channel后sender携程会自己关闭
	m.MessageSenderWaitGroup.Wait()
	wg.Done()
}
