package main

import "fmt"

var pint1, pint2 int                                   // two package variables with default value of int on package level
var pbool1, pstring1, pfloat1 = true, "pstring1", 3.14 // when we use initial values, we can ommit the type

func main() {
	fmt.Println(pint1)
	fmt.Println(pint2)
	fmt.Println(pbool1)
	fmt.Println(pstring1)
	fmt.Println(pfloat1)

	var fint1, fint2 int                                    // two variables on function level with default value of int on package level
	var fbool1, fstring1, ffloat1 = false, "fstring1", 6.66 // when we use initial values, we can ommit the type
	fmt.Println(fint1)
	fmt.Println(fint2)
	fmt.Println(fbool1)
	fmt.Println(fstring1)
	fmt.Println(ffloat1)

	fstring2 := fstring1 // Short variable declaration inside functions
	fmt.Println(fstring2)
}
