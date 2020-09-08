package main

import "fmt"

func main() {

	/*
		evaluate first call later
	*/
	x := 1
	addX := func(a int) int {
		x = x + a
		fmt.Println("Evaluate first!")
		return x
	}
	defer fmt.Println(addX(3) + 5)
	fmt.Println(x)
}
