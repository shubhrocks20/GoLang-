package main

import "fmt"

const age = 30
var gender = 0
func main() {
	const name string = "golang"
	fmt.Println(name)

	fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)
}