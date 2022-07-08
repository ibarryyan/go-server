package test

import (
	"count_num/pkg/utils"
	"fmt"
	"testing"
)

func TestS(t *testing.T) {
	str := utils.NowTimeStamp()
	fmt.Println(str)

	timeStr := utils.TimeStamp2NowTimeStr(str)
	fmt.Println(timeStr)

	stamp := utils.NowTimeStr2TimeStamp(timeStr)
	fmt.Println(stamp)
}
