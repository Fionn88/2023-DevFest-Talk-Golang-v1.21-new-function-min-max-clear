package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Before clear:", slice)

	clear(slice)

	fmt.Println("After clear:", slice)

	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Before clear:", myMap)

	clear(myMap)

	fmt.Println("After clear:", myMap)
}
