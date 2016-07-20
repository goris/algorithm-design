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
	pivot := Partition(arr, start, end)
	if DEBUG {
		fmt.Println("arr: ", arr, "start: ", start, "end: ", end, "pivot: ", pivot)
	}
	if start < pivot-1 {
		QuickSort(arr, start, pivot-1)
	}
	if pivot < end {
		QuickSort(arr, pivot, end)
	}
	fmt.Println("Comps: ", comps)
}

func Partition(arr []int, left int, right int) int {
	pivot := arr[left]
	i := left + 1

	for j := left + 1; j < right; j++ {
		comps++
		if arr[j] < pivot {
			swap(arr, i, j)
			i++
		}
	}
	swap(arr, left, i-1)
	return i
}
