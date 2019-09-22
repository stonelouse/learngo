package main

import (
	"fmt"
	mathutil "github.com/stonelouse/learngo/tour02functions/mathutils"
	"github.com/stonelouse/learngo/tour02functions/stringutil"
)

// not exported function with two parameters
func sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Printf("5 + 2 = %d\n", sum(5, 2))
	fmt.Printf("5 * 2 = %d\n", mathutil.Mul(5, 2))

	a, b := mathutil.Swap(5, 2)
	fmt.Printf("Swap(5, 2) = %d, %d\n", a, b)

	fmt.Println(stringutil.MySplitIn2("First part, second part", ","))
}
