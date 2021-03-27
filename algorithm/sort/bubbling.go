package main

import "fmt"

func main() {
	arr := []int{1, 4, 5, 6, 73, 11, 2}
	BubblingSort(arr)
	fmt.Println(arr)
}
func BubblingSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
