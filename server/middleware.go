package server

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/sirupsen/logrus"
	"time"
)

var ginLogger = logrus.WithField("server", "gin")
var wechatMsgLogger = logrus.WithField("server", "wechatMsg")

// ginRequestLog gin 请求纪录( 请求状态码 处理时间 请求方法 IP 路由 )
func ginRequestLog(c *gin.Context) {
	// 开始时间
	startTime := time.Now()

	// 处理请求
	c.Next()

	// 结束时间
	endTime := time.Now()

	// 执行时间
	latencyTime := endTime.Sub(startTime)

	// 请求方式
	reqMethod := c.Request.Method

	// 请求路由
	reqUri := c.Request.RequestURI

	// 状态码
	statusCode := c.Writer.Status()

	// 请求IP
	clientIP := c.ClientIP()

	// 日志格式
	ginLogger.Infof(
		"| %3d | %13v | %15s | %s | %s |",
		statusCode, latencyTime, clientIP, reqMethod, reqUri,
	)
}

func wechatMsgLog(m *Message) {
	// 开始时间
	startTime := time.Now()

	// 处理请求
	m.Next()

	// 结束时间
	endTime := time.Now()

	// 执行时间
	latencyTime := endTime.Sub(startTime)

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(m.Reply)
	reply := buffer.String()

	var msgType, key, id string
	if m.MsgType == message.MsgTypeEvent {
		msgType = string(m.Event)
		key = m.EventKey
	} else {
		msgType = string(m.MsgType)
		key = m.Content
	}
	if m.UnionID != "" {
		id = m.UnionID
	} else {
		id = string(m.FromUserName)
	}
	wechatMsgLogger.Infof(
		"| %10v | %v | %v | %s | %s |",
		latencyTime, msgType, key, id, reply)
}
