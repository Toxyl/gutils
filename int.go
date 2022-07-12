package gutils

import (
	"math/rand"
	"strconv"
	"time"
)

func StringToInt64(s string, defaultValue int64) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i = defaultValue
	}
	return i
}

func StringToInt32(s string, defaultValue int32) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		i = int64(defaultValue)
	}
	return int32(i)
}

func StringToInt(s string, defaultValue int) int {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		i = int64(defaultValue)
	}
	return int(i)
}

func BytesToInt(b []byte, defaultValue int) int {
	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		i = int64(defaultValue)
	}
	return int(i)
}

func GetRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	n := max - min + 1
	if n <= 0 {
		return min
	}
	return rand.Intn(n) + min
}
