package main

import (
	"fmt"
	"math"
)

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func apply(a, b float64, op func(float64, float64) float64) (float64, error) {
	if op == nil {
		return math.NaN(), fmt.Errorf("apply: nil operation")
	}
	return op(a, b), nil
}

func main() {
	a, _ := apply(1, 2, add) // a := add(1, 2)
	b, _ := apply(1, 2, sub) // b := sub(1, 2)
	c, _ := apply(1, 2, nil)
	fmt.Println(a, b, c)
}
