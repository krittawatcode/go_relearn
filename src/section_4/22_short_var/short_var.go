package main

import "fmt"

func main() {
	a, b := 1, 2
	c, b := b, a
	fmt.Println(a, b, c)
}
