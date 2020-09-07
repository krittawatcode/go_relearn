package main

import "fmt"

func createF() func() int {
	var x = 0
	return func() int {
		x++
		return x
	}
}

func main() {

	f := createF()
	// fmt.Println(f)
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3

	f2 := createF()
	fmt.Println(f2()) // 1
	fmt.Println(f2()) // 2
	fmt.Println(f2()) // 3
}
