package gutils

import "strings"

func GetLastError(err error) string {
	sl := strings.Split(err.Error(), ":")
	s := sl[len(sl)-1]
	return strings.TrimSpace(s)
}
