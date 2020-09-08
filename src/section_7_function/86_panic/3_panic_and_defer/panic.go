package main

import (
	"fmt"
	"os"
)

var cygwinPath = os.Getenv("CYGWIN")

func A(x int) {
	defer func() {
		fmt.Println("Defer in A")
	}()
	B(x)
}

func B(x int) {
	defer func() {
		fmt.Println("Defer in B")
	}()
	C(x)
}

func C(x int) {
	defer func() {
		fmt.Println("Defer in C")
	}()
	if x == 4 {
		panic("for no reason")
	}
}

func main() {
	A(4)
}
