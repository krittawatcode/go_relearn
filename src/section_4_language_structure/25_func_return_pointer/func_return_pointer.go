package main

import "fmt"

func fp() *int {
	x := 4
	return &x
}

func main() {
	x := fp()
	fmt.Printf("%T, %d\n", x, *x)
	fmt.Println(x)
	fmt.Println(x == fp()) // false because this fp() doesn't same as fp() upper!! <- create new local var everytimes we call func
}
