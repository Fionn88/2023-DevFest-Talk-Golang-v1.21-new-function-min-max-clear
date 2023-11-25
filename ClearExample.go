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

	// The Original Method
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Before clear:", s)
	// s = s[:0]
	s = nil
	fmt.Println("After clear:", s)

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Before clear:", m)
	// for k := range m {
	// 	delete(m, k)
	// }
	m = nil
	fmt.Println("After clear:", m)

	s := []int{1, 2, 3, 4, 5}
	// Approach 1
	s = s[:0]
	// Approach 2
	s = nil

	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// Approach 1
	for k := range m {
		delete(m, k)
	}
	// Approach 2
	m = nil
}
