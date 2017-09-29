package log

import (
	"log"
	"os"
	"runtime"
)

type LevelMsg struct {
	level   int
	message string
}

//提供异步输出的web服务log打印
type WebAppender struct {
	BaseAppender
	messages chan *LevelMsg
}

const (
	_defaultChanSize = 10000
)

func newWebAppender() *WebAppender {
	var appender = &WebAppender{}
	appender.messages = make(chan *LevelMsg, _defaultChanSize)
	appender.logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//default to all level
	appender.levels = make(map[int]bool)
	for _, l := range Levels {
		appender.levels[l] = true
	}

	go appender.flushLog()
	return appender
}

func (this *WebAppender) Output(msg string, level int) {
	this.messages <- &LevelMsg{level, msg}
}
func (this *WebAppender) Close() {
	//TODO
	close(this.messages)
}

func (this *WebAppender) flushLog() {
	var lm *LevelMsg
	for {
		select {
		case lm = <-this.messages:
			this.writeLog(lm.message, lm.level)
		}
	}

}
func (this *WebAppender) writeLog(msg string, level int) {
	if ok := this.levels[level]; !ok {
		return
	}
	if goos := runtime.GOOS; goos == "windows" {
		this.logger.Println(msg)
	} else {
		this.logger.Println(colorBrushes[level](msg))
	}
}
