package main

import "fmt"

func main() {
	fruits := [5]string{
		"apple",
		"banana",
		"papaya",
		"orange",
		"mango",
	}
	// slices - use concept of pointer -> mean that if slice change then array change
	myFavor := fruits[1:4] // start at 1 end exclude 4
	fmt.Println(len(myFavor), cap(myFavor), myFavor)
	yourFavor := myFavor
	yourFavor[0] = "guava" // if you change dup slice it will affect all slice and array
	fmt.Println(len(yourFavor), cap(yourFavor), yourFavor)
	fmt.Println(len(myFavor), cap(myFavor), myFavor)
	fmt.Println(len(fruits), cap(fruits), fruits)
}
