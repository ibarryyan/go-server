package utils

import (
	"time"
)

// NowTimeStr return 2022-01-21 12:21:31
func NowTimeStr() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

// NowTimeStamp return 1657255820
func NowTimeStamp() int64 {
	return time.Now().Unix()
}

// TimeStamp2NowTimeStr  1657255820 -> 2022-01-21 12:21:31
func TimeStamp2NowTimeStr(stamp int64) string {
	format := time.Unix(stamp, 0).Format("2006-01-02 15:04:05")
	return format
}

// NowTimeStr2TimeStamp  2022-01-21 12:21:31 -> 1657255820
func NowTimeStr2TimeStamp(str string) int64 {
	var LOC, _ = time.LoadLocation("Asia/Shanghai")
	tim, _ := time.ParseInLocation("2006-01-02 15:04:05", str, LOC)
	return tim.Unix()
}
