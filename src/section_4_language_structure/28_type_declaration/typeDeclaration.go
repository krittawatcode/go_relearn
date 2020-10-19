package main

import "fmt"

type KG float64
type LB float64

func main() {
	// a := 2
	// A := 3
	// fmt.Println(a + A) // 5

	// b := "2"
	// B := "3"
	// fmt.Println(b + B) // 23

	/*
		type declalation

		type-yourTypeName-underlyingType

		ex. type KG float64
			type LB float64
	*/

	k := KG(3)
	lb := LB(3)

	fmt.Println(k == 3.0) // A constant may be given a type explicitly by a constant declaration or conversion
	result := k + 3       // An untyped constant has a default type which is the type to which the constant is implicitly converted in contexts where a typed value is required
	fmt.Println(result)
	fmt.Printf("%T\n", result)

	fmt.Println(k == KG(lb)) // cast underlying type
}
