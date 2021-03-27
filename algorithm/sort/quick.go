package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 1, 4, 5, 6, 73, 11, 2}
	// QuickSort(arr)
	QuickSort(arr)

	fmt.Println(arr)
}
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[0]
	left, right := 0, len(arr)-1

	for left < right {
		for arr[right] >= pivot && left < right {
			right--
		}

		arr[left] = arr[right]

		for arr[left] <= pivot && left < right {
			left++
		}

		arr[right] = arr[left]

	}
	arr[left] = pivot
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}
