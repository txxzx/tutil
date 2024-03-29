package tutil

import "time"

/**
    @date: 2022/9/20
**/

// GetAgeByBirthDayTS 跟据生日时间戳获取年龄
func GetAgeByBirthDayTS(ts int64) int {
	if ts <= 0 {
		return 0
	}
	// 当前时间
	nowTm := time.Now()
	// 生日时间
	birthTm := time.Unix(ts, 0)
	subYear := nowTm.Year() - birthTm.Year()

	if subYear <= 0 {
		return 0
	}
	// 已经满虚岁了，+1
	if nowTm.YearDay() > birthTm.YearDay() {
		return subYear + 1
	}
	return subYear
}
