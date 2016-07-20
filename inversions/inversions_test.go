package inversions

import (
	"fmt"
	"testing"
)

func TestCountInversionsO2Ordered(t *testing.T) {

	fmt.Println("Testing CountInversionsO2Ordered")
	var arr []int
	arr = []int{1, 2, 3, 4, 5}
	expected := 0
	res := CountInversionsO2(arr)

	if res != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, res)
	}
}

func TestCountInversionsO2Small(t *testing.T) {

	fmt.Println("Testing CountInversionsO2Small")
	var arr []int
	arr = []int{1, 20, 6, 4, 5}
	expected := 5
	res := CountInversionsO2(arr)

	if res != expected {
		t.Errorf("expected\n%s\n,got:\n%s", expected, res)
	}
}
