package main

import "fmt"

func intervene() func() {
	fmt.Println("1")
	return func() {
		fmt.Println("2")
	}
}

func main() {
	defer intervene()() // evaluate first return later
	fmt.Println("*******")
}
