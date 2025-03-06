package tools

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Logger *logrus.Entry

func NewLoggeer() {
	logStore := logrus.New()
	logStore.SetLevel(logrus.DebugLevel)

	// 同时写到多个输出
	w1 := os.Stdout                                                               //写到控制台
	w2, _ := os.OpenFile("./vote.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644) //赋值权限 读 创造 添加
	logStore.SetOutput(io.MultiWriter(w1, w2))                                    // io.MultiWriter 返回一个 io.Writer 对象

	logStore.SetFormatter(&logrus.JSONFormatter{}) //json 形式
	Logger = logStore.WithFields(logrus.Fields{
		"name": "香香编程喵喵喵",
		"app":  "voteV2",
	})
	//可以增加hook 函数，当触发某些特殊的日志后，执行某些函数。比如：邮件告警，日志分割与上报等。
	//Logger.AddHook()

	//context 上下文
	//logStore.WithContext()
}
