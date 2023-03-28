package roman_convert

import (
	"bytes"
)

/**
    @date: 2023/3/28
**/

func IntToRoman(num int) string {
	numMap := map[int]string{
		1000: "M", 900: "CM", 500: "D", 400: "CD",
		100: "C", 90: "XC", 50: "L", 40: "XL",
		10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I",
	}

	var buffer bytes.Buffer
	for num >= 1000 {
		// 写到string缓冲里面
		buffer.WriteString("M")
		num = num - 1000
	}
	index := 100
	for num > 0 {
		if num/index > 0 {
			if value, ok := numMap[num/index*index]; ok {
				buffer.WriteString(value)
			} else if num/index > 5 {
				buffer.WriteString(numMap[5*index])
				for temp := (num/index - 5) * index; temp > 0; temp = temp - index {
					buffer.WriteString(numMap[index])
				}
			} else if num/index < 5 {
				for temp := num / index; temp > 0; temp-- {
					buffer.WriteString(numMap[index])
				}
			}
		}
		num = num - (num/index)*index
		index = index / 10
	}
	return buffer.String()
}

func RomanToInt2(s string) int {
	if s == "" {
		return 0
	}
	dict := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var out int
	for i := 0; i < len(s); i++ {   // MMCMIX
		if i < len(s)-1 && dict[rune(s[i])] < dict[rune(s[i+1])] {
			// 获取value值
			out -= dict[rune(s[i])]
		} else {
			out += dict[rune(s[i])]
		}
	}
	return out
}