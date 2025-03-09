package main

import "fmt"

// variadic function -> n no of parameters acceptance
func sum(nums ...int) int {
	total := 0
	for _,num := range nums {
		total += num
	}
	return total
}
func main() {
	nums := []int{3,4,5,6}
	res := sum(nums...)
	fmt.Println(res)
	
}