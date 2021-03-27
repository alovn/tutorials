package main

import (
	"fmt"
	"sort"
)

func main() {
	k := Constructor(3, []int{})
	fmt.Println(k.Add(3))
	fmt.Println(k.Add(5))
	fmt.Println(k.Add(10))
	// fmt.Println(k.Add(9))
	// fmt.Println(k.Add(4))
}

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {

	this := KthLargest{k: k, nums: nums}
	this.sort()
	if len(this.nums) >= this.k {
		this.nums = this.nums[len(this.nums)-this.k:]
	}
	return this
}

func (this *KthLargest) Add(val int) int {
	if len(this.nums) >= this.k {
		if val < this.nums[0] {
			return this.nums[0]
		}
		this.nums[0] = val
		this.sort()
		return this.nums[0]
	} else {
		this.nums = append([]int{val}, this.nums...)
		this.sort()
		if len(this.nums) == this.k {
			return this.nums[0]
		}
		return 0
	}
}

func (this *KthLargest) sort() {
	sort.Stable(sort.IntSlice(this.nums))
}
