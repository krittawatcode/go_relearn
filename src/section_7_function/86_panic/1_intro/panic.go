package main

import (
	"fmt"
	"os"
)

func main() {
	x := []int{1, 2, 3}
	// fmt.Println(x[10]) // panic out of range

	/* ----- part 1 -----
	//	fmt.Println(x[0])
	//	panic("for no reason")
	//	fmt.Println(x[2])
	*/

	fmt.Println(x[1])
	f, err := os.Open("sasf")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(x[20])
	fmt.Println(f)

}
