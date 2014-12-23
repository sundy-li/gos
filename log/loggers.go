package log

import (
	"fmt"
)

//log containner
var (
	_appenders = map[string]Appender{}
	_loggers   = map[string]*Logger{}
)

func init() {
	_loggers["root"] = rootLog
}

func Get(name string) *Logger {
	if _log, ok := _loggers[name]; ok {
		return _log
	}
	return rootLog
}

func Trace(msg string) {
	rootLog.WriteMsg(msg, LevelTrace, _fileLevel)
}

func Info(msg string) {
	rootLog.WriteMsg(msg, LevelInfo, _fileLevel)
}

func Warn(msg string) {
	rootLog.WriteMsg(msg, LevelWarn, _fileLevel)
}

func Debug(msg string) {
	rootLog.WriteMsg(msg, LevelDebug, _fileLevel)
}

func Error(msg string) {
	rootLog.WriteMsg(msg, LevelError, _fileLevel)
}

func Tracef(msg string, args ...interface{}) {
	rootLog.WriteMsgFunc(msg, LevelTrace, _fileLevel, fmt.Sprintf, args...)
}

func Infof(msg string, args ...interface{}) {
	rootLog.WriteMsgFunc(msg, LevelInfo, _fileLevel, fmt.Sprintf, args...)
}

func Warnf(msg string, args ...interface{}) {
	rootLog.WriteMsgFunc(msg, LevelWarn, _fileLevel, fmt.Sprintf, args...)
}
func Debugf(msg string, args ...interface{}) {
	rootLog.WriteMsgFunc(msg, LevelDebug, _fileLevel, fmt.Sprintf, args...)
}
func Errorf(msg string, args ...interface{}) {
	rootLog.WriteMsgFunc(msg, LevelError, _fileLevel, fmt.Sprintf, args...)
}
