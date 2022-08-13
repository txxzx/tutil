package tutil

import (
	"encoding/json"
	"strings"
)

/**
    @date: 2022/8/13
**/

// map转为string
func MapToString(intMap map[int]int) (string, error) {
	bt, err := json.Marshal(intMap)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

// string 转为 map
func StringToMap(str string) (intmap map[int]int) {
	mapMath := make(map[int]int)
	err := json.Unmarshal([]byte(str), &mapMath)
	if err != nil {
		return nil
	}
	return mapMath
}

// 字符串切片转为数字切片
func StringSliceToIntSlice(str []string) []string {
	st := strings.Join(str, ",")
	s := strings.Split(st, ",")
	return s
}

// 切片转为字符串
func SLiceToString(s []int) string {
	b, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(b)
}
