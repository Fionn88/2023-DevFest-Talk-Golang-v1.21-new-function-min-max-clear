package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Before clear:", slice) // Before clear: [1 2 3 4 5]

	clear(slice)

	fmt.Println("After clear:", slice) // After clear: [0 0 0 0 0]

	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Before clear:", myMap) // Before clear: map[a:1 b:2 c:3]

	clear(myMap)

	fmt.Println("After clear:", myMap) // After clear: map[]

	// The Original Method
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Before clear:", s) // Before clear: [1 2 3 4 5]
	// s = s[:0]
	s = nil
	fmt.Println("After clear:", s) // After clear: []

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Before clear:", m) // Before clear: map[a:1 b:2 c:3]

	// for k := range m {
	// 	delete(m, k)
	// }
	m = nil
	fmt.Println("After clear:", m) // After clear: map[]
}
