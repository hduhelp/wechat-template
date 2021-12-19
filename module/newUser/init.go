// 新关注自动回复模组
// TODO: INIT Welcome Message

package newUser

import (
	"github.com/hduhelp/wechat-template/server"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"sync"
)

func init() {
	instance = &newUser{}
	server.RegisterModule(instance)
}

var instance *newUser

type newUser struct {
	WelcomeMessage string
}

func (newUser) GetModuleInfo() server.ModuleInfo {
	return server.ModuleInfo{
		ID:       server.NewModuleID("atom", "newUser"),
		Instance: instance,
	}
}

func (newUser) Init() {

}

func (newUser) PostInit() {
}

func (newUser) Serve(s *server.Server) {
	s.MsgEngine.EventSubscribe(0, func(msg *server.Message) {
		text := message.NewText("Hi!很高兴见到你。")
		msg.Reply = &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})
}

func (newUser) Start(server *server.Server) {

}

func (newUser) Stop(server *server.Server, wg *sync.WaitGroup) {
	defer wg.Done()
}
