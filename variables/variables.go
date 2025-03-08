package main

import "fmt"

func main() {
	// string declaration
	// var name string = "golang"
	// infer the type
	var name = "golang"
	var isAdult = true
	fmt.Println("Am i Adult: ", isAdult)

	var age int = 22
	fmt.Println("my name is: ",name)
	fmt.Println("my age is : ", age)


	// shorthand syntax
	nickname := "golang"
	fmt.Println(nickname)

	var team string

	team = "india"

	fmt.Println("my team is : ", team)

	// var price float32 = 50.5
	var price = 50.5
	// price := 50.5

	fmt.Println(price)
}