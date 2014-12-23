package utils

import (
	"testing"
)

func TestCurFile(t *testing.T) {
	f, l := CurFileLine(0)
	println(f, l)

	ff, ll := CurFileLine(0)
	println(ff, ll)
}
