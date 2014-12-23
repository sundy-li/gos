package log

import (
	"log"
	"reflect"
)

const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

var (
	AppenderTypeFile    = "file"
	AppenderTypeConsole = "console"
)

var (
	Levels    = []int{LevelTrace, LevelInfo, LevelWarn, LevelDebug, LevelError}
	LevelMsgs = map[int]string{
		LevelTrace: "Trace",
		LevelInfo:  "Info",
		LevelWarn:  "Warn",
		LevelDebug: "Debug",
		LevelError: "Error",
	}
)

type Appender interface {
	Output(msg string, level int)
	Close()
}

type BaseAppender struct {
	Appender
	levels map[int]bool
	logger *log.Logger
}

func getLevelFromStr(level string) int {
	var a = LevelTrace
	switch level {
	case "trace":
		a = LevelTrace
	case "info":
		a = LevelInfo
	case "warn":
		a = LevelWarn
	case "debug":
		a = LevelDebug
	case "error":
		a = LevelError
	default:
		log.Printf("error level %s, default to trace level \n", level)
	}
	return a
}
func (appender *BaseAppender) SetLevel(levels interface{}) {
	t := reflect.TypeOf(levels)
	v := reflect.ValueOf(levels)
	appender.levels = make(map[int]bool)
	switch t.Kind() {
	case reflect.Int:
		v := levels.(string)
		lv := getLevelFromStr(v)
		for _, l := range Levels {
			if l < lv {
				continue
			}
			appender.levels[l] = true
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			lv := getLevelFromStr(v.Index(i).Interface().(string))
			appender.levels[lv] = true
		}
	default:
		log.Fatalf("set levels error: %#v\n", levels)
	}
}
