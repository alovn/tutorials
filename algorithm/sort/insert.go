package main

import "fmt"

func main() {
	arr := []int{1, 4, 5, 6, 73, 11, 2}
	InsertSort(arr)
	fmt.Println(arr)
}
func InsertSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
