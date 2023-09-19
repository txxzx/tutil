package sort

import "fmt"

/**
    @date: 2023/9/19
**/

/*
 1.选择排序，每次从待排序的数组中选择一个最小的放到数组的开始位置
*/
// SelectSort
func SelectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 遍历数组
	for i := 0; i < len(arr); i++ {
		// 定义一个最小的元素
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		// 每次找到一个最小的元素，进行交换位置
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

//选择排序
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	selectsort := SelectSort(arr)
	fmt.Println(selectsort)
}

// 获取最大元素
func SelectMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	}

	max := arr[0]
	for i := 0; i < length; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
