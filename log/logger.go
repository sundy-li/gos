package log

import (
	"fmt"
	"github.com/sundy-li/gos/utils"
)

var (
	rootLog      = NewDefaultLogger()
	rootAppender = newConsoleAppender()
	_fileLevel   = 2
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

func (this *Logger) Trace(msg string) {
	this.WriteMsg(msg, LevelTrace, _fileLevel)
}

func (this *Logger) Info(msg string) {
	this.WriteMsg(msg, LevelInfo, _fileLevel)
}

func (this *Logger) Warn(msg string) {
	this.WriteMsg(msg, LevelWarn, _fileLevel)
}

func (this *Logger) Debug(msg string) {
	this.WriteMsg(msg, LevelDebug, _fileLevel)
}

func (this *Logger) Error(msg string) {
	this.WriteMsg(msg, LevelError, _fileLevel)
}

func (this *Logger) Tracef(msg string, args ...interface{}) {
	this.WriteMsgFunc(msg, LevelTrace, _fileLevel, fmt.Sprintf, args...)
}

func (this *Logger) Infof(msg string, args ...interface{}) {
	this.WriteMsgFunc(msg, LevelInfo, _fileLevel, fmt.Sprintf, args...)
}

func (this *Logger) Warnf(msg string, args ...interface{}) {
	this.WriteMsgFunc(msg, LevelWarn, _fileLevel, fmt.Sprintf, args...)
}
func (this *Logger) Debugf(msg string, args ...interface{}) {
	this.WriteMsgFunc(msg, LevelDebug, _fileLevel, fmt.Sprintf, args...)
}
func (this *Logger) Errorf(msg string, args ...interface{}) {
	this.WriteMsgFunc(msg, LevelError, _fileLevel, fmt.Sprintf, args...)
}
