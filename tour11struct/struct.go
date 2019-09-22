package main

import "fmt"

// a struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	// Struct literals
	var v Vertex = Vertex{1, 2}
	var (
		v2 = Vertex{X: 3}
		v3 = Vertex{}
		p4 = &Vertex{Y: 4}
	)
	fmt.Println(v)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(*p4) // {0 4}
	fmt.Println(p4)  // &{0 4}

	// struct fields are accessed using dot
	fmt.Printf("v=[%d, %d]\n", v.X, v.Y)

	// Its not necessary to denote an explicit derefernce when using struct pointer
	p := &v
	fmt.Printf("v=<%d, %d>\n", (*p).X, p.Y)
}
