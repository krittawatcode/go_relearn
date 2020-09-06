package main

import (
	"fmt"
	"reflect"
)

func avg(a, b float64) float64 {
	return (a + b) / 2.0
}

func avg2(a, b float64, c int, d, e, f int32) (float64, string, int) {
	return 4.4, "gg", 555
}

func avg3(_, _ float64, _ int, _, _, _ int32) (float64, string, int) {
	return 4.4, "gg", 555
}

func main() {
	// a := (1 + 3) / 2.0
	a := avg(1, 3)
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(avg))
	fmt.Println(reflect.TypeOf(avg2))
}
