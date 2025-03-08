package main

import "fmt"

func main() {
	age := 10

	// if age >= 18 {
	// 	fmt.Println("Person is an Adult!")
	// } else {
		// 	fmt.Println("Person is not an Adult!")
		// }
		
	if age >= 18 {
		fmt.Println("Person is an Adult!")
	} else if age >= 12 {
		fmt.Println("Person is tennager")
	} else{
		fmt.Println("Person is a kid")
	}

	var role = "admin"
	var hasPermission = false

	if role == "admin" && hasPermission {
		fmt.Println("yes")
	}
	// we can declare a variable inside if construct
	if age := 15; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Person is tennager", age)
	}
	// go doesnot have ternary operator, you will have to use normal if else

	
}