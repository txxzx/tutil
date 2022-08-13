package tutil

/**
    @date: 2022/8/13
**/

// 判断输入三个,输出一个确定数字是否在两个数字之间,返回boole值
func CheckMemeberIs(min, max, num int) bool {
	return num >= min && num <= max
}
