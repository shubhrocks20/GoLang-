package main

import "fmt"

// numbered sequence of specific length
func main() {
	// zero values
	var nums [4]int
	// array length
	fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums)

	var newArray [4]string
	newArray[1] = "shubham"
	fmt.Println(newArray)
	// to declare in single line
	numbers := [3]int{1,2,3}
	fmt.Println(numbers)

	// 2d arrays

	numbers2 := [2][2]int{{1,2},{3,4}}
	fmt.Println(numbers2)


	// - fixed size, that is predictable
	// - Memory optimization
	// - Constant time access

}