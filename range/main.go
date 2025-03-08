package main

import "fmt"

// range -> iterating over data structure
func main() {
	// nums := []int {6,7,8}
	// sum := 0
	// for idx, num := range nums {
	// 	fmt.Println(idx)
	// 	sum += num
	// }
	// fmt.Println(sum)

	// m := map[string]string{"fname":"john", "lname":"doe"}
	// for k,v := range m{
	// 	fmt.Println(k, v)
	// }
	// unicode code
	// starting byte of rune
	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}