package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Before clear:", slice)
	// Result: Before clear: [1 2 3 4 5]

	clear(slice)

	fmt.Println("After clear:", slice)
	// Result: After clear: [0 0 0 0 0]

	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Before clear:", myMap)
	// Result: Before clear: map[a:1 b:2 c:3]

	clear(myMap)

	fmt.Println("After clear:", myMap)
	// Result: After clear: map[]
}
