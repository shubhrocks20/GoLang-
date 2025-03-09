package main

import "fmt"

// func add(a int, b int) (int) {
// 	return a + b
// }

// func getLanguages()(string,string,string) {
// 	return "golang","C","java"
// }

// func processIt(fn func(a int) int) {
// 	fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}

}
func main() {
	// fn := func (a int ) int{
	// 	return a
	// }
	fn := processIt()
	fmt.Println(fn(5))
	// res := add(1, 2)
	// fmt.Println(res)

	// fmt.Println(getLanguages())

}