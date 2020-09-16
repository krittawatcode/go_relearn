package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := -5
	fmt.Printf("%b\n", x)
	fmt.Printf(strconv.FormatInt(5, 2)) // change 5 base 10 to 5 base 2
}
