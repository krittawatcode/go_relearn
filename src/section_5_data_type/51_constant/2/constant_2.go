package main

import (
	"fmt"
	"reflect"
)

type Peice int

func main() {

	const persons = 4
	toffee := 5 / persons
	cost := 2.0 / persons

	fmt.Println(toffee, reflect.TypeOf(toffee))
	fmt.Println(cost, reflect.TypeOf(cost))

	// untpyed constant
	toffeeV2 := Peice(5) / persons
	fmt.Println(toffeeV2, reflect.TypeOf(toffeeV2))

	// declair group constant
	const (
		Sun = 1
		Mon = 2
		Tue = 3
	)
	fmt.Println(Sun, Mon, Tue)

	const (
		Wed = iota + 4 // ex. float64 = iota + 1.1
		Thu
		Fri // int
		Sat
	)
	fmt.Println(Wed, Thu, Fri, Sat)

}
