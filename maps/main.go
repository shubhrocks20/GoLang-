package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
	// creating map
	m := make(map[string]string)
	
	// setting an element
	m["name"] = "shubham"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["name"], m["area"])
	fmt.Println(m["avd"]) // if key doesnot exist in the map then it returns zero value

	newmap := make(map[int]int)
	newmap[1] = 1
	fmt.Println(newmap[1]) // 1
	fmt.Println(newmap[101]) // 0 as 101 not in map

	fmt.Println(len(m))

	// delete(m, "area")
	clear(m)
	fmt.Println(m)


	// fmt.Println(k)
	
	// k := map[string]int{"price":40, "phone": 3}

	// li, ok := k["phones"]
	// fmt.Println(li)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok", )
	// }

	c1 := map[string]int{"price":30, "phone": 2}
	c2 := map[string]int{"price":30, "phone": 9}

	fmt.Println(maps.Equal(c1, c2))
	
}