package log

import (
	"fmt"
	"strings"
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

func Trace(msgs ...string) {
	msg := strings.Join(msgs, " ")
	rootLog.WriteMsg(msg, LevelTrace, _defaultLevel)
}

func Info(msgs ...string) {
	msg := strings.Join(msgs, " ")
	rootLog.WriteMsg(msg, LevelInfo, _defaultLevel)
}

func Warn(msgs ...string) {
	msg := strings.Join(msgs, " ")
	rootLog.WriteMsg(msg, LevelWarn, _defaultLevel)
}

func Debug(msgs ...string) {
	msg := strings.Join(msgs, " ")
	rootLog.WriteMsg(msg, LevelDebug, _defaultLevel)
}

func Error(msgs ...string) {
	msg := strings.Join(msgs, " ")
	rootLog.WriteMsg(msg, LevelError, _defaultLevel)
}

func Tracef(msg string, args ...interface{}) {
	rootLog.WriteMsgFunc(msg, LevelTrace, _defaultLevel, fmt.Sprintf, args...)
}

func Infof(msg string, args ...interface{}) {
	rootLog.WriteMsgFunc(msg, LevelInfo, _defaultLevel, fmt.Sprintf, args...)
}

func Warnf(msg string, args ...interface{}) {
	rootLog.WriteMsgFunc(msg, LevelWarn, _defaultLevel, fmt.Sprintf, args...)
}
func Debugf(msg string, args ...interface{}) {
	rootLog.WriteMsgFunc(msg, LevelDebug, _defaultLevel, fmt.Sprintf, args...)
}
func Errorf(msg string, args ...interface{}) {
	rootLog.WriteMsgFunc(msg, LevelError, _defaultLevel, fmt.Sprintf, args...)
}
