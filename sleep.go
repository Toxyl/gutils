package gutils

import "time"

func SleepSeconds(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}

func RandomSleep(min, max int, duration time.Duration) {
	time.Sleep(time.Duration(GetRandomInt(min, max)) * duration)
}
