package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	x := complex128(complex(7, 3))
	y := complex128(complex(1, 4))
	fmt.Println(x + y)

	fmt.Println(cmplx.Sqrt(-9))
	fmt.Println(cmplx.Abs(complex(3, -4)))
}
