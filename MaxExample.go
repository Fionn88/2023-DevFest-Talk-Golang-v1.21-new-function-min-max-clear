package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {

	// math.Max math.Mix Function
	fmt.Println(math.Max(20, 10))          // Result: 20
	fmt.Println(math.Max(10, 10))          // Result: 10
	fmt.Println(math.Max(-10, 20))         // Result: 20
	fmt.Println(math.Max(10, -20))         // Result: 10
	fmt.Println(math.Max(-10, -10))        // Result: -10
	fmt.Println(math.Max(10.5, -20.12))    // Result: 10.5
	fmt.Println(math.Max(math.Inf(2), 10)) // Result: +Inf
	fmt.Println(math.Max(10, math.NaN()))  // Result: Nan

	fmt.Println(math.Min(0.0, -0.0)) // Result: 0
	fmt.Println(math.Min(20, 10))    // Result: 10

	// max function
	fmt.Println(max(math.Inf(1), 2, 3.0, math.Inf(-1))) // Result: +Inf
	fmt.Println(max("one", "two", "three"))             // Result: two

	// min function
	fmt.Println(min(math.Inf(-1), 2, 3.0, math.Inf(-1))) // Result: -Inf
	fmt.Println(min("one", "two", "three"))              // Result: one

	// slices.Min
	x := []float64{2, -5, 8, 1.2}
	fmt.Println(slices.Min(x)) // Result: -5
	fmt.Println(slices.Max(x)) // Result: 8

}
