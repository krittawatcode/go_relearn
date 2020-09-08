package main

import "fmt"

func main() {

	/*
		"defer" will called after every func
		called by "down to top" if we have many defer

		defer is like "stack" by using "defer stack" -> first in last out

		evaluate first, call later (see in part 2)
	*/

	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	fmt.Println("hello")

}
