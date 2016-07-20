package quicksort

import (
	"fmt"
	"github.com/goris/algorithm-design/quicksort"
	"testing"
)

func TestQuickSort(t *testing.T) {

	fmt.Println("Testing QuickSort")
	var arr []int
	var expected []int
	arr = []int{3, 9, 8, 4, 6, 10, 2, 5, 7, 1}
	expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	quicksort.QuickSort(arr, 0, len(arr))

	if arr[9] != expected[9] {
		t.Errorf("expected\n%v\n,got:\n%v", expected, arr)
	}
}

/* func TestCountQuickSortComparisonsLeft(t *testing.T) {

    fmt.Println("Testing QuickSort")
    var arr []int
    var expected []int
    arr = []int{3,9,8,4,6,10,2,5,7,1}
    expected = []int{1,2,3,4,5,6,7,8,9,10}
    quicksort.QuickSort(arr, 0, len(arr)-1)

	if arr[9] != expected[9] {
		t.Errorf("expected\n%v\n,got:\n%v", expected, arr)
	}
}
*/
