package main

import "fmt"

func curry(a int) func(fn func(int) int) int {
	return func(fn func(int) int) int {
		return fn(a) // Closure with parameter a!
	}
}

func main() {
	fn := curry(2)
	fmt.Printf("%d\n", fn(pow)) // 4

	fmt.Printf("%d\n", fn(func(x int) int { return x * x * x })) // 8
}

func pow(x int) int {
	return x * x
}
