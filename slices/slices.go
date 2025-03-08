package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// + useful methods
func main() {
	// uninitialized slice is nill
	// var nums []int
	// // fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 2, 5)
	// // capacity -> maximum numbers of elements can fit
	// fmt.Println(cap(nums))
	// // fmt.Println(nums)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// fmt.Println(nums)
	// fmt.Println(cap(nums)) // 20

	// nums2 := []int{}
	// nums2 = append(nums2, 1)
	// nums2[0] = 101
	// fmt.Println(nums2)
	// fmt.Println(len(nums2))
	// fmt.Println(cap(nums2))
	
	// // copy function
	// var a = make([]int, 0, 5)
	// a = append(a, 2)
	// var b = make([]int, len(a))
	// copy(b, a)
	// fmt.Println(a, b)


	// slice operator
	var nums = []int{1,2,3}
	fmt.Println((nums[:1])) // a,b b is exclusive

	// slice 
	var nums1 = []int{1,2,3}
	var nums2 = []int{1,2,4}

	fmt.Println(slices.Equal(nums1, nums2))

	var nums3 = [][]int{{1,2,3}, {4,5,6}}
	fmt.Println(nums3)

}
