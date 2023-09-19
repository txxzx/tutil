package sort

import (
	"fmt"
	"testing"
)

/**
    @date: 2023/9/19
**/

func TestSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	selectsort := SelectSort(arr)
	fmt.Println(selectsort)
}

func TestQuickSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(QuickSort(arr))
}

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(GetMax(arr))
	fmt.Println(BubbleSort(arr))
}
