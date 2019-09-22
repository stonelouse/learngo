package main

import (
	"fmt"
	"strings"
)

func basics() {
	/*
		A slice is a
		... dynamically-sized, flexible view
		... into the elements of an array.
		A slice does not store any data, it just describes a section of an underlying array.
	*/
	//               0  1  2  3   4   5
	primes := [6]int{2, 3, 5, 7, 11, 13}

	//  The type []T is a slice with elements of type T.
	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// a[low : high] includes the low-st element but exclude the high-st element
	var slice1 []int = primes[1:4] // [3 5 7] !!!!!!!!!!!!!!
	fmt.Println(slice1)

	slice2 := primes[2:6]
	slice1[2] = 17
	fmt.Println(primes) // [2 3 5 17 11 13]
	fmt.Println(slice2) // [5 17 11 13]

	colorSlice := []string{"red", "blue", "green"} // this creates an array and then builds a slice that reference it
	fmt.Println(colorSlice)                        // [red blue green]

}

func lengthAndCapacity() {
	primes := [6]int{2, 3, 5, 17, 11, 13}

	// The default is zero for the low bound and the length of the slice for the high bound.
	slice3 := primes[:]
	fmt.Println(slice3) // [2 3 5 17 11 13]
	slice4 := primes[3:]
	fmt.Println(slice4) // [17 11 13]
	slice5 := primes[:len(primes)/2]
	fmt.Println(slice5) // [2 3 5]

	/*
		A slice has a length
		... the length of a slice is the number of elements it contains.

		A slice has a capacity.
		... the capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

		The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	*/
	fmt.Printf("len(primes)=%d cap(primes)=%d\n", len(primes), cap(primes)) // len(primes)=6 cap(primes)=6
	fmt.Printf("len(slice4)=%d cap(slice4)=%d\n", len(slice4), cap(slice4)) // 	len(slice4)=3 cap(slice4)=3
	fmt.Printf("len(slice5)=%d cap(slice5)=%d\n", len(slice5), cap(slice5)) // 	len(slice5)=3 cap(slice5)=6

	// You can extend a slice's length by re-slicing it, ...
	slice5 = slice5[:cap(slice5)]
	fmt.Printf("len(slice5)=%d cap(slice5)=%d slice5=%v\n", len(slice5), cap(slice5), slice5) // len(slice5)=6 cap(slice5)=6 slice5=[2 3 5 17 11 13]
	// ... provided it has sufficient capacity.
	// slice5 = slice5[:cap(slice5)+1]
	// fmt.Printf("len(slice5)=%d cap(slice5)=%d slice5=%v\n", len(slice5), cap(slice5), slice5) // panic: runtime error: slice bounds out of range [:7] with capacity 6
}

func nilSlices() {
	/*
		The zero value of a slice is nil.

		A nil slice has a length and capacity of 0 and has no underlying array.
	*/
	var slice []int
	fmt.Println(slice, len(slice), cap(slice), slice == nil) // [] 0 0 true
}

func createSliceWithMake() {
	/*
		Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	*/
	slice := make([]int, 5)                    // allocates a zeroed array and returns a slice with length and capacity 5
	fmt.Println(slice, len(slice), cap(slice)) // [0 0 0 0 0] 5 5

	slice = make([]int, 5, 8)                  // allocates a zeroed array and returns a slice with length 5 and capacity 8
	fmt.Println(slice, len(slice), cap(slice)) // [0 0 0 0 0] 5 8
	slice = slice[:cap(slice)]
	slice[7] = 99
	fmt.Println(slice, len(slice), cap(slice)) // [0 0 0 0 0 0 0 99] 8 8
}

func slicesOfslices() {
	board := [][]string{
		[]string{"__", "__", "__"},
		[]string{"__", "__", "__"},
		[]string{"__", "__", "__"},
	}
	board[0][0] = "1A"
	board[2][2] = "2A"
	board[1][2] = "1B"
	board[1][0] = "2B"
	board[0][2] = "1C"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func appendingToSlices() {
	var s []int // nil slice
	printSlice(s)

	// Go provides a built-in append to append new elements to a slice
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed
	s = append(s, 1)
	printSlice(s)

	// We can append a variable number of elements
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	basics()
	lengthAndCapacity()
	nilSlices()
	createSliceWithMake()
	slicesOfslices()
	appendingToSlices()
}
