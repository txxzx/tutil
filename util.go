package tutil

/**
    @date: 2022/8/10
**/

import (
	"fmt"
	mathRand "math/rand"
	"strings"
	"time"
)

/*
从数组里面随机取出一个数
*/
func RandMember(rand []int32) int32 {
	// 没有随机数种子会出现每次随机的数都是一样的
	r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	// 根据随机的数取出数据
	rint := r.Intn(len(rand))
	member := rand[rint]
	// 获取对应的随机数的个数 会随机到0
	member = int32(r.Intn(20) + 1)
	return member
}

/**
 *
 * @param x int整型
 * @return int整型
 */
func reverse(x int) int {
	// write code here
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}
	// 判断是否溢出
	if res > (1<<31)-1 || res < -(1<<31)-1 {
		return 0
	}
	return res
}

// 数字转换成大写
func Transfer(num int) string {
	chineseMap := []string{"", "十", "百", "千", "万", "十", "百", "千", "亿", "十", "百", "千"}
	chineseNum := []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	listNum := []int{}
	for ; num > 0; num = num / 10 {
		listNum = append(listNum, num%10)
	}

	fmt.Println(listNum)

	n := len(listNum)
	chinese := ""
	//注意这里是倒序的

	for i := n - 1; i >= 0; i-- {
		fmt.Println(chineseNum[listNum[i]])
		fmt.Println(chineseMap[i])
		chinese = fmt.Sprintf("%s%s%s", chinese, chineseNum[listNum[i]], chineseMap[i])
	}

	fmt.Println(chinese)
	//注意替换顺序
	for {
		copychinese := chinese
		copychinese = strings.Replace(copychinese, "零万", "万", 1)
		copychinese = strings.Replace(copychinese, "零亿", "亿", 1)
		copychinese = strings.Replace(copychinese, "零十", "零", 1)
		copychinese = strings.Replace(copychinese, "零百", "零", 1)
		copychinese = strings.Replace(copychinese, "零千", "零", 1)
		copychinese = strings.Replace(copychinese, "零零", "零", 1)
		copychinese = strings.Replace(copychinese, "零圆", "圆", 1)

		if copychinese == chinese {
			break
		} else {
			chinese = copychinese
		}
	}

	chinese = strings.Replace(chinese, "零", "", -1)
	chinese = strings.Replace(chinese, "一十", "十", -1)
	return chinese
}
