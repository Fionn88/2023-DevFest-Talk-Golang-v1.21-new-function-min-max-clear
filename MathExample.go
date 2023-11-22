package main

import (
	"fmt"
	"math"
)

func main() {

	// math.Max math.Mix Function
	fmt.Println(math.Max(20, 10))
	fmt.Println(math.Max(10, 10))
	fmt.Println(math.Max(-10, 20))
	fmt.Println(math.Max(10, -20))
	fmt.Println(math.Max(-10, -10))
	fmt.Println(math.Max(10.5, -20.12))
	fmt.Println(math.Min(0.0, -0.0))
	fmt.Println(math.Min(20, 10))

	fmt.Println(math.Max(math.Inf(2), 10))
	fmt.Println(math.Max(10, math.NaN()))

	// Output:
	// 20
	// 10
	// 20
	// 10
	// -10
	// 10.5
	// 0
	// 10
	// +Inf
	// NaN

	// max function
	fmt.Println(max(math.Inf(1), 2, 3.0, math.Inf(-1)))
	fmt.Println(max("one", "two", "three"))

	// Output:
	// +Inf
	// two

	// min function
	fmt.Println(min(math.Inf(-1), 2, 3.0, math.Inf(-1)))
	fmt.Println(min("one", "two", "three"))

	// Output:
	// -Inf
	// one

}
