package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	makeAMap()
	createAMapLiterally()
	accessMap()
}

func makeAMap() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{ // struct literal
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}

func createAMapLiterally() {
	m := map[string]Vertex{ // map literal
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google Inc": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Google Inc"])

	for k, v := range m {
		fmt.Printf("%s--%v\n", k, v)
	}
}

func accessMap() {
	m := make(map[string]int)

	m["Answer"] = 41 // insert element
	fmt.Println(m)
	m["Answer"] = 42 // update element
	fmt.Println(m)

	fmt.Printf("%d\n", m["Answer"])  // retrieve element
	fmt.Printf("%d\n", m["nowhere"]) // retrieve element: 0
	v, ok := m["Answer"]
	fmt.Printf("%d ... %v\n", v, ok)
	v, ok = m["nowhere"]
	fmt.Printf("%d ... %v\n", v, ok)

	delete(m, "Answer")
	fmt.Println(m)
}
