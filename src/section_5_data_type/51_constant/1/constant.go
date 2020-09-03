package main

import "fmt"

func main() {
	var x = 2
	var y = 0
	const z = 0 // evaluated value when declair <- complie time
	// debugging will skip line 8 because z init at complie time (debug = runtime with stop)
	fmt.Println(x / y) // can run because it will calculate at runtime -> will get erroe at runtime
	// fmt.Println(x / z) // can't run because z will get value at compile time not runtime
}
