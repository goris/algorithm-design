package inversions

import (
)

func CountInversionsO2(arr []int)(int) {
    var count int

    for i := 0; i < len(arr) - 1 ; i++ {
        for j := i + 1; j < len(arr); j++ {
            if arr[i] > arr[j] {
                count++;
            }
        }
    }

    return count

}
