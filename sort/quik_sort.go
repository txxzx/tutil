package sort

/**
    @date: 2023/9/19
**/

/*
// 快速排序
 1.找个基准元素
 2.快速排序是先找个基准元素，然后左边元素合基准元素进行比较，进行移动，
	比基准元素大的移动到基准元素的右边。如果比基准元素小，就移动到基准元素的左边。每次都会向中间移动一个元素
 3.采取递归进行处理
*/

// QuickSort
func QuickSort(arr []int) []int {
	// 递归推出的条件
	if len(arr) <= 1 {
		return arr
	}

	spliteData := arr[0]
	// 比我小的数组
	lowArr := make([]int, 0)
	// 表我大的数组
	heightArr := make([]int, 0)
	// 和我一样大的数组
	midArr := make([]int, 0)

	// 把我加儒中间元素里面
	midArr = append(midArr, spliteData)
	for i := 1; i < len(arr); i++ {
		if arr[i] < spliteData {
			lowArr = append(lowArr, arr[i])
		} else if arr[i] > spliteData {
			heightArr = append(heightArr, arr[i])
		} else {
			midArr = append(midArr, arr[i])
		}
	}

	// 递归的进行处理元素
	lowArr, heightArr = QuickSort(lowArr), QuickSort(heightArr)
	finishedArr := append(append(lowArr, midArr...), heightArr...)
	return finishedArr
}
