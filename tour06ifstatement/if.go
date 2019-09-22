package main

import "fmt"
import "math"

func main() {
	fmt.Printf("2 is odd? %v\n", isOdd(2))
	fmt.Printf("3 is odd? %v\n", isOdd(3))
	fmt.Printf("2.55 is odd? %v\n", isOddAfterFloor(2.55))
	fmt.Printf("3.44 is odd? %v\n", isOddAfterFloor(3.44))
}

func isOdd(value int) bool {
	if value%2 == 0 {
		return false
	}
	return true
}

func isOddAfterFloor(value float64) bool {
	t := math.Floor(value)
	fmt.Printf("...%T...%v\n", t, t)

	/*
		the if statement can start with a short statement to execute before the condition!
		... variables declared inside an if short statement are available inside the if block and any else block
	*/
	if v := int(math.Floor(value)); v%2 == 0 {
		return false
	}
	return true
}
