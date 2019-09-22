package main

import "fmt"

/*
  A defer statement defers the execution of a function until the surrounding function returns.
  The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

  Deferred function calls are pushed onto a stack.
  When a function returns, its deferred calls are executed in last-in-first-out order.
*/

var greet = "Reaper"

func main() {
	defer fmt.Printf("... %s", greet)
	greet = "Angel"
	fmt.Println("Follow the")

	fmt.Println("\ncounting")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d\n", i)
	}
	fmt.Println("done")
}
