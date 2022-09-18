package tutil

import (
	"fmt"
	"strconv"
)

/**
    @date: 2022/9/18
**/

func StringToInt32(s string) int32 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int32(i)
}

func Int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

// int32数组转化为string数组
func Int32ToStringSlice(a []int32) []string {
	r := make([]string, len(a))
	for k, v := range a {
		r[k] = Int32ToString(v)
	}
	return r
}

func StringToInt32Slice(a []string) []int32 {
	t := make([]int32, len(a))
	for k, v := range a {
		t[k] = StringToInt32(v)
	}
	return t
}

//StringSliceToInt64Slice
func StringSliceToInt64Slice(str []string) (ids []int64) {
	for _, i := range str {
		id, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	return
}

// int64转为string
func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

//Folat32ToInt64
func FloatToInt64(count float32) (int64, error) {
	countStr := fmt.Sprintf("%0.0f", count)
	return strconv.ParseInt(countStr, 10, 64)
}

func InterfaceToString(inte interface{}) (str string) {
	rs := ""
	switch inte.(type) {
	case int, int32, int64:
		rs = fmt.Sprintf("%d", inte)
	case float32, float64:
		rs = fmt.Sprintf("%0.f", inte)
	case string:
		rs = fmt.Sprintf("%s", inte)
	}
	return rs
}
