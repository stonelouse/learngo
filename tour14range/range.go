package main

import "fmt"

func main() {
	rangeFormOfFor()
	omitIndexOrCopy()
}

func rangeFormOfFor() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	// The range form of the for loop iterates over a slice or map.
	/*
		When ranging over a slice,
		... two values are returned for each iteration.
		... The first is the index, and
		... the second is a COPY of the element at that index.
	*/
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func omitIndexOrCopy() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
