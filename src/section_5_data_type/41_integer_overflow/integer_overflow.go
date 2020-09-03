package main

import "fmt"

func main() {
	x := int8(127)
	y := int8(1)
	X := 127
	Y := 2
	fmt.Println(x + y)
	fmt.Println(X + Y)
	fmt.Println(int8(X + Y))
	fmt.Printf("%b", 129)

	a := 3.999
	fmt.Println(int(a))
}
