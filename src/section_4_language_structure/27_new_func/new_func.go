package main

import "fmt"

func newIntPointer() *int {
	var x int
	return &x
}

func main() {
	// x := newIntPointer()
	fmt.Println(newIntPointer() == newIntPointer())
	fmt.Println(new(int) == new(int)) // build-in new int pointer -< return pointer of allowcate type
}
