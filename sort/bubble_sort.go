package sort

/**
    @date: 2023/9/19
**/

/*
冒泡排序：每次从待冒泡的数组里面两两进行比较然后交换位置，冒到最后就是最大一个元素
*/

// BubbleSort
func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 冒泡的逻辑
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// 冒泡选取最大元素
func GetMax(arr []int) int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	return arr[len(arr)-1]
}
