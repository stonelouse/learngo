package main

import "fmt"

func main() {
	// The type [n]T is an array of n values of type T.
	// An array's length is part of its type, so arrays cannot be resized.
	var a [10]int

	a[0] = 'a'
	fmt.Printf("%c\n", a[0])
	fmt.Println(a) // [97 0 0 0 0 0 0 0 0 0]

	// Array literals
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	primes2 := []int{2, 3, 5} // this creates an array and then builds a slice that reference it
	fmt.Println(primes2)
}
