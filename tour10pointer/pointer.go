package main

import "fmt"

func main() {
	var i int = 42
	var p *int = &i // p points to i

	fmt.Printf("*p=%d\n", *p) // read i through the pointer p
	*p = 666                  // set i through the pointer p
	fmt.Printf("i=%d\n", i)

	j := 7411
	p = &j // p points now to j
	fmt.Printf("*p=%d\n", *p)
	*p += 1
	fmt.Printf("j=%d\n", j)
}
