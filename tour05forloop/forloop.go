package main

import "fmt"

func main() {
	var sum int = 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("Sum=%v\n", sum)

	sum = 1

	for sum < 100 { // init and post statement are optional
		sum += sum
		fmt.Printf("...=%v\n", sum)
	}
	fmt.Printf("Sum=%v", sum)
}
