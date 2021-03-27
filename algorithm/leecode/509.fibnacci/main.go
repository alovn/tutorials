package main

import "fmt"

func main() {
	//0, 1, 1, 2, 3, 5, 8
	fmt.Println(fib(5))
	fmt.Println(fib2(5))
	fmt.Println(fib3(5))
}

func fib2(n int) int {
	var result, tmp int = 0, 1
	for i := 0; i < n; i++ {
		result += tmp
		tmp = result - tmp
	}
	return result
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	var last1, last2 int = 0, 1
	var result int
	for i := 1; i < n; i++ {
		result = last1 + last2
		last1 = last2
		last2 = result
	}
	return result
}

func fib3(n int) int {
	if n < 2 {
		return n
	}
	return fib3(n-2) + fib(n-1)
}
