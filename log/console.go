package log

import (
	"log"
	"os"
	"runtime"
)

var (
	colorBrushes = map[int]Brush{
		LevelTrace: NewBrush("1;36"), // Trace      cyan
		LevelInfo:  NewBrush("1;32"), // Info       green
		LevelDebug: NewBrush("1;34"), // Debug      blue
		LevelWarn:  NewBrush("1;33"), // Warn       yellow
		LevelError: NewBrush("1;31"), // Error      red
	}
)

type Brush func(string) string

func NewBrush(color string) Brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}

type ConsoleAppender struct {
	BaseAppender
}

func (this *ConsoleAppender) Output(msg string, level int) {
	if ok := this.levels[level]; !ok {
		return
	}

	if goos := runtime.GOOS; goos == "windows" {
		this.logger.Println(msg)
	} else {
		this.logger.Println(colorBrushes[level](msg))
	}
}
func (this *ConsoleAppender) Close() {

}

func newConsoleAppender() *ConsoleAppender {
	console := &ConsoleAppender{}
	console.logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//default to all level
	console.levels = make(map[int]bool)
	for _, l := range Levels {
		console.levels[l] = true
	}
	return console
}
