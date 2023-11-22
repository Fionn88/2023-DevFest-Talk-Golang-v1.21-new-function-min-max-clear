package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Before clear:", slice)
	// Before clear: [1 2 3 4 5]

	clear(slice)

	fmt.Println("After clear:", slice)
	// After clear: [0 0 0 0 0]

	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Before clear:", myMap)
	// Before clear: map[a:1 b:2 c:3]

	clear(myMap)

	fmt.Println("After clear:", myMap)
	// After clear: map[]
}
