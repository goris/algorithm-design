package main

import (
	"fmt"
	_ "github.com/goris/algorithm-design/inversions"
	"github.com/goris/algorithm-design/quicksort"
	_ "github.com/goris/algorithm-design/reader"
	_ "log"
)

func main() {
	/*
		lines, err := reader.ReadInts("data/integersInversions.txt")
		if err != nil {
			log.Fatalf("readInts: %s", err)
		}
		inversions := inversions.CountInversionsO2(lines)
		fmt.Println("# of inversions: ", inversions)
	*/

	arr := []int{3, 8, 2, 5, 1, 4, 7, 6}
	fmt.Println("arr: ", arr)
	quicksort.QuickSort(arr, 0, len(arr)-1)
	fmt.Println("Res arr: ", arr)

}
