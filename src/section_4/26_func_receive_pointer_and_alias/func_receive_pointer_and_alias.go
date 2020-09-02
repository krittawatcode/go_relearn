package main

import "fmt"

func double(x *int) {
	*x = *x * 2
}

func main() {
	x := 2
	double(&x)
	fmt.Println(x)
}
