package main

import "fmt"

func create() (fs []func()) {
	for i := 0; i < 2; i++ {
		fs = append(fs, func() {
			fmt.Println(i)
		})
	}
	return
}
