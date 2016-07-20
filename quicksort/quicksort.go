package quicksort

import (
	"fmt"
	_ "os"
)

const DEBUG = true

var comps int

func swap(arr []int, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
	if DEBUG {
		//fmt.Println("After swap: ", arr)
	}
}

func QuickSort(arr []int, start, end int) {
	if start < end {
		pivot := Partition(arr, start, end)
		//pivot := arr[start]
		if DEBUG {
			fmt.Println("arr: ", arr, "start: ", start, "end: ", end, "pivot: ", pivot)
		}
		QuickSort(arr, start, pivot-1)
		QuickSort(arr, pivot+1, end)
	}
	fmt.Println("Comps: ", comps)
}

func Partition(arr []int, left int, right int) int {
	pivot := arr[right]
	i := left

	for j := left; j <= right-1; j++ {
		if arr[j] <= pivot {
			swap(arr, i, j)
			i++
		}
	}
	swap(arr, i, right)
	return i
}
