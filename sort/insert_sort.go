package sort

/**
    @date: 2023/9/19
**/

// standardInsertSort 插入排序
func StandardInsertSort(list []int) []int {
	resultList := make([]int, 0)
	length := len(list)
	i := 1
	resultList = append(resultList, list[0])
	for i < length {
		for j := 0; j < len(resultList); j++ {
			if list[i] <= resultList[j] {
				resultList = insertList(resultList, j, list[i])
				break
			}
			if j == len(resultList)-1 && list[i] > resultList[j] {
				resultList = insertList(resultList, j+1, list[i])
				break
			}
		}
		i++
	}
	return resultList
}
func insertList(list []int, i int, x int) []int {
	returnList := []int{}
	n := 0
	if i == len(list) {
		returnList = append(list, x)
		return returnList
	}
	for n < len(list) {
		if n < i {
			returnList = append(returnList, list[n])
		} else if n == i {
			returnList = append(returnList, x)
			returnList = append(returnList, list[n])
		} else {
			returnList = append(returnList, list[n])
		}
		n++
	}
	return returnList
}
