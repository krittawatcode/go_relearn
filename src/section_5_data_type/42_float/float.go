package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64 // we cant do this if type of x is int
	fmt.Println(x, -x, 1/x, 1/-x)
	fmt.Println(math.IsNaN(x / x))
	fmt.Println(math.IsInf(1/x, 0))
}
