package main

import (
	"fmt"
	"math"
)

func sqrt(x float64, s float64) (float64, int) {
	/*
		given a number x,
		... we want t ofind the number z for which z * z is most nearly z
		... next guess z = z - ( z*z - x) / (2*z)

		if x = 4 ...
		1st guess 2.5 = 2.5  - (2.5  * 2.5  - 4) / (2 * 2.5) = 2.5 - 2.25 / 5 = 2.5 - 0.45 = 2.05
		2nd guess     = 2.05 - (2.05 * 2.05 - 4) / (2 * 2.05)                              = 2.0006098

	*/
	z := s
	i := 0
	goodEnough := isGoodEnough(x, z)
	for ; !goodEnough; i++ {
		z -= (z*z - x) / 2 * z
		goodEnough = isGoodEnough(x, z)
		if i > 100000 {
			break
		}
	}
	return z, i
}

func isGoodEnough(x, z float64) bool {
	q := z * z
	diff := math.Abs(q - x)
	result := diff < 0.01
	fmt.Printf("...x=%v ...z=%v ...q=%v ...diff=%v ...result=%v\n", x, z, q, diff, result)
	return result
}

func main() {
	result, iterations := sqrt(4.0, 2.0)
	fmt.Printf("Result=%v in Iteratation=%v\n", result, iterations)
	result, iterations = sqrt(2.0, 1.0) // 1.414213562373095
	fmt.Printf("Result=%v in Iteratation=%v\n", result, iterations)
}
