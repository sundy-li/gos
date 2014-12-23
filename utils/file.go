package utils

import (
	"runtime"
	"strings"
)

func CurFileLine(addLevel int) (filePath string, line int) {
	var _level int
	var filename string
	for i := 0; i < 20; i++ {
		_, filename, _, _ = runtime.Caller(i)
		if strings.HasSuffix(filename, "gos/utils/file.go") {
			_level = i + 1
			break
		}
	}
	_, filePath, line, _ = runtime.Caller(_level + addLevel)
	return
}

func GetFileName(filePath string) string {
	var paths = strings.Split(filePath, "/")
	return paths[len(paths)-1]
}
