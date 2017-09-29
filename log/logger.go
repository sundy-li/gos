package log

import (
	"fmt"
	"strings"

	"github.com/sundy-li/gos/utils"
)

var (
	rootLog       = NewDefaultLogger()
	rootAppender  = newConsoleAppender()
	_defaultLevel = 2
)

type Logger struct {
	appenders []Appender

	LoggerConfig
}

type Msg struct {
	Message string
	Level   int
}

func NewDefaultLogger() *Logger {
	l := &Logger{
		appenders: []Appender{
			rootAppender,
		},
	}
	l.ShowFileLine = true
	return l
}

func (this *Logger) WriteMsg(msg string, level int, fileLevel int) {
	if this.Level > level {
		return
	}
	levelMsg := LevelMsgs[level]

	//func call
	if this.ShowFileLine {
		fileName, line := utils.CurFileLine(fileLevel)
		if this.ShortFile {
			fileName = utils.GetFileName(fileName)
		}
		msg = fmt.Sprintf("%s:%d %s", fileName, line, msg)
	}
	msg = "[" + levelMsg + "] " + msg

	for _, appender := range this.appenders {
		appender.Output(msg, level)
	}
}

func (this *Logger) WriteMsgFunc(msg string, level int, fileLevel int, f func(str string, args ...interface{}) string, args ...interface{}) {
	this.WriteMsg(f(msg, args...), level, fileLevel+1)
}

func (this *Logger) Trace(msgs ...string) {
	msg := strings.Join(msgs, " ")
	this.WriteMsg(msg, LevelTrace, _defaultLevel)
}

func (this *Logger) Info(msgs ...string) {
	msg := strings.Join(msgs, " ")
	this.WriteMsg(msg, LevelInfo, _defaultLevel)
}

func (this *Logger) Warn(msgs ...string) {
	msg := strings.Join(msgs, " ")
	this.WriteMsg(msg, LevelWarn, _defaultLevel)
}

func (this *Logger) Debug(msgs ...string) {
	msg := strings.Join(msgs, " ")
	this.WriteMsg(msg, LevelDebug, _defaultLevel)
}

func (this *Logger) Error(msgs ...string) {
	msg := strings.Join(msgs, " ")
	this.WriteMsg(msg, LevelError, _defaultLevel)
}

func (this *Logger) Tracef(msg string, args ...interface{}) {
	this.WriteMsgFunc(msg, LevelTrace, _defaultLevel, fmt.Sprintf, args...)
}

func (this *Logger) Infof(msg string, args ...interface{}) {
	this.WriteMsgFunc(msg, LevelInfo, _defaultLevel, fmt.Sprintf, args...)
}

func (this *Logger) Warnf(msg string, args ...interface{}) {
	this.WriteMsgFunc(msg, LevelWarn, _defaultLevel, fmt.Sprintf, args...)
}
func (this *Logger) Debugf(msg string, args ...interface{}) {
	this.WriteMsgFunc(msg, LevelDebug, _defaultLevel, fmt.Sprintf, args...)
}
func (this *Logger) Errorf(msg string, args ...interface{}) {
	this.WriteMsgFunc(msg, LevelError, _defaultLevel, fmt.Sprintf, args...)
}
