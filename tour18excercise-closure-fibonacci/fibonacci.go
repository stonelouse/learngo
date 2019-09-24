package main

import "fmt"

/*
	F0 = 0 Def.
	F1 = 1 Def.
	F2 = 1 Def.
	F3 = 2
	F4 = 3
	F5 = 5
	F6 = 8
	...
	Fn = Fn-1 + Fn-2 for n >= 3
*/
func fibonacciGenerator() func() int {
	var (
		last       int = 1
		beforeLast int = 1
		next       int = 0
	)
	return func() int {
		var result int = 0

		switch next {
		case 0:
			result = 0
		case 1:
			result = 1
		case 2:
			result = 1
		default:
			result = beforeLast + last
			beforeLast = last
			last = result
		}
		next++

		return result
	}
}

func main() {
	f := fibonacciGenerator()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", f())
	}
}
