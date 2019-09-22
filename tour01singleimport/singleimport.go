package main

// paranthesized, "factored" import statement
import (
	"fmt"
	"math"
)

/*
import "fmt"
import "math"
*/
// works also; but is autoformatted by VSCode extension

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

// VSCode: Ctrl+Shift+P: ´Go: Install current package´
// %GOPATH%\bin\tour01singleimport
