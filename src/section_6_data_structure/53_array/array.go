package main

import (
	"fmt"
	"reflect"
)

func main() {

	fruits := [5]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"mango",
	}

	fmt.Println(len(fruits), fruits)

	twoSlots := [2]int{}
	threeSlots := [3]int{}

	fmt.Println("Two slot\t", reflect.TypeOf(twoSlots))
	fmt.Println("Three slot\t", reflect.TypeOf(threeSlots))
	/*
		it's mean type of array in go depend on size
	*/

	// ellipsls
	fruits2 := [...]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"mango",
	}
	fmt.Println(len(fruits2), fruits2)

	// set index
	fruits3 := [...]string{
		4: "apple",
		2: "banana",
		1: "papaya",
		3: "orange",
		0: "mango",
	}
	fmt.Println(len(fruits3), fruits3, fruits3[4])

	fruits4 := [...]string{
		4: "apple",
		"banana",
		"papaya",
		"orange",
		"mango",
	}
	fmt.Println(len(fruits4), fruits4, fruits4[2]) // fruits4[2] = empty string

	// compare array
	fmt.Println([2]int{1, 2} == [2]int{1, 2})
	fmt.Println([3]int{1, 2, 3} == [3]int{1, 3, 2})

	// arrays copy by value
	dubFruits := fruits
	dubFruits[0] = "Watermelon"
	fmt.Println(fruits)
	fmt.Println(dubFruits)

	// use pointer = same ref
	ptrFruits := &fruits
	ptrFruits[0] = "Strawberry"
	fmt.Println(fruits)
	fmt.Println(*ptrFruits)

	// Array[][]
	a := [2][3]int{{1, 2, 3}, {3, 4, 5}}
	fmt.Println(a)

}
