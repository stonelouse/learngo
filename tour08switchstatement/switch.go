package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var win = "windows"

	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds. http://127.0.0.1:3999/flowcontrol/10
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.") // no break necessary
	case win: // no constant necessary
		fmt.Println("... windows")
	default:
		fmt.Printf("%s. \n", os)
	}

	alternativeLongIfElseChain()
}

func alternativeLongIfElseChain() {
	t := time.Now()

	switch { // without condition same as switch true
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
