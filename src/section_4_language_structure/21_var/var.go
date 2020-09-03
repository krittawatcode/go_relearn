package main

import "fmt"

func main() {
	// var
	var x int = 1 + 2 // if not initiate value then it will be zero velue of that type
	var y = 4
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%#v\n", x)

	//--------------------
	var z []string
	fmt.Println(z == nil)
	fmt.Printf("%T\n", z)

}
