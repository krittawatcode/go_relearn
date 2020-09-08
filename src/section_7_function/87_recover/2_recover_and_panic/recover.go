package main

import (
	"fmt"
	"os"
)

var cygwinPath = os.Getenv("CYGWIN")

func A(x int) {
	defer func() {
		fmt.Println("Defer in A")
		if p := recover(); p != nil {
			panic(fmt.Sprintf("a: %v", p))
		}
	}()
	B(x)
	fmt.Println("Hello in A")
}

func B(x int) {
	defer func() {
		fmt.Println("Defer in B")
		if p := recover(); p != nil {
			panic(fmt.Sprintf("b: %v", p))
		}
	}()
	C(x)
	fmt.Println("Hello in B")
}

func C(x int) {
	defer func() {
		fmt.Println("Defer in C")
		if p := recover(); p != nil {
			panic(fmt.Sprintf("c: %v", p))
		}
	}()
	if x == 4 {
		panic("for no reason")
	}
}

func main() {
	A(4)
}
