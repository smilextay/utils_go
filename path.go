package utils

import (
	"path/filepath"
	"strings"
)

func Parent(dir string) string {

	dir, _ = filepath.Abs(dir)
	index := strings.LastIndex(dir, "/")
	if index == -1 {
		return dir
	}
	return substr(dir, 0, index)
}
func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
