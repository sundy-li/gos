package utils

import (
	"runtime"
	"strings"
)

func CurFileLine(addLevel int) (filePath string, line int) {
	_, filePath, line, _ = runtime.Caller(addLevel)
	//regardless gopath
	index := strings.Index(filePath, "src/")
	if index != -1 {
		filePath = filePath[index+4:]
	}
	return
}

func GetFileName(filePath string) string {
	var paths = strings.Split(filePath, "/")
	return paths[len(paths)-1]
}
