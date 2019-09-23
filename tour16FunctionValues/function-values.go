package main

import (
	"fmt"
	"math"
)

/*
	Function values

	Go supports Higher order functions.

	Functions are values too.

	Function values may be used as function arguments and return values.
	They can be passed around just like other values.
*/
func compute(fn func(float64, float64) float64) float64 {
	return fn(5, 12)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
