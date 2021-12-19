package pong

import (
	"github.com/gin-gonic/gin"
	"github.com/hduhelp/wechat-template/server"
	"sync"
)

func init() {
	instance = &pong{}
	server.RegisterModule(instance)
}

var instance *pong

type pong struct {
}

// GetModuleInfo 获取模块信息
// 返回值为 `server.ModuleInfo` 类型，其中包含模块 ID 和一个 module 实例指针.
// 一般照抄，修改模块名称和相应字符串即可
func (m *pong) GetModuleInfo() server.ModuleInfo {
	return server.ModuleInfo{
		ID:       server.NewModuleID("atom", "pong"),
		Instance: instance,
	}
}

func (m *pong) Init() {
	// 初始化过程
	// 在此处可以进行 Module 的初始化配置
	// 如配置读取
}

func (m *pong) PostInit() {
	// 第二次初始化
	// 再次过程中可以进行跨Module的动作
	// 如通用数据库等等
}

func (m *pong) Serve(server *server.Server) {
	// 注册服务函数部分
	server.HttpEngine.GET("/ping", handlePingPong)
}

func (m *pong) Start(server *server.Server) {
	// 此函数会新开携程进行调用
	// ```go
	// 		go exampleModule.Start()
	// ```

	// 可以利用此部分进行后台操作
	// 如http服务器等等
}

func (m *pong) Stop(server *server.Server, wg *sync.WaitGroup) {
	// 别忘了解锁
	defer wg.Done()
	// 结束部分
	// 一般调用此函数时，程序接收到 os.Interrupt 信号
	// 即将退出
	// 在此处应该释放相应的资源或者对状态进行保存
}

func handlePingPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg":        "pong",
		"User-Agent": c.GetHeader("User-Agent"),
	})
}
