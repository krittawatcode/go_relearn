package main

import "fmt"

func main() {
	x := "hi-สวัสดี" // []byte base16, immutable
	y := []byte{0x68, 0x69, 0x2d, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}
	z := []rune(x)
	// --------	h	, i	  , -   , ส               , ว               ,
	// ------------------------- should be
	fmt.Println(string(y[0])) // h
	fmt.Println(string(y[1])) // i
	fmt.Println(string(y[2])) // -
	fmt.Println(string(y[3])) // ส
	fmt.Println(string(y[3])) // ว

	fmt.Println(len(x))
	fmt.Println(len(y))
	fmt.Println(len(z))
}
