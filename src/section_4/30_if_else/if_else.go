package main

import "fmt"

func main() {
	x := 5

	if x < 2 {
		fmt.Println("if")
	} else {
		fmt.Println("else")
	}

	// --------------------

	if v := x + 1; (x <= 5) || false {
		fmt.Println("if ", v)
	} else {
		fmt.Println("else")
	}
}
