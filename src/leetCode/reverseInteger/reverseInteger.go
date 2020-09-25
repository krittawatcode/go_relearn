package main

import (
	"fmt"
	"math"
)

func main() {
	x := reverse(123)
	fmt.Println(x)
	y := 123
	y = y / 10 * 10
	fmt.Println(y)
}

func reverse(x int) (out int) {
	for ; x != 0; x /= 10 {
		out = out*10 + x%10
		if out > math.MaxInt32 || out < -math.MaxInt32-1 {
			return 0
		}
	}
	return out
}
