package main

import "fmt"

// by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

// by reference
func changeNum(num *int){
	*num = 5
	fmt.Println("In ChangeNum: ", *num)
}
func main() {
	num := 1
	changeNum(&num)
	// fmt.Println("memory address: ", &num)
	fmt.Println("Global Num:", num)

}