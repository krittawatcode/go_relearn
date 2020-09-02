package main

import "fmt"

func main() {
	x := 70
	var p *int // pointer of int
	p = &x
	fmt.Println("&x", &x)
	fmt.Println("*&x", *&x)
	fmt.Println("p", p)
	fmt.Println("&p", &p)
	fmt.Println("*p", *p)
}
