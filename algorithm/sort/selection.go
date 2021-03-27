package main

import "fmt"

func main() {
	arr := []int{1, 4, 5, 6, 73, 11, 2}
	SelectionSort(arr)
	fmt.Println(arr)
}
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
