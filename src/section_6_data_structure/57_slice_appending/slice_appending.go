package main

import "fmt"

func main() {
	x := [10]int{}
	a := x[0:5]
	b := x[5:7]
	for i := range a { // size of a -> len(a)
		a[i] = i * i
	}
	b[0] = 98
	b[1] = 99
	// ----- Part I -----
	// a = append(a, b...)
	// a[len(a)-1] = 25
	// fmt.Println("x: \t", x)
	// fmt.Println("a: \t", len(a), cap(a), a)
	// fmt.Println("b: \t", len(b), cap(b), b)

	// a = append(a, b...)
	// fmt.Println("a: \t", len(a), cap(a), a)

	a = append(a, b...)
	a[len(a)-1] = 25
	a = append(a, 71, 72, 73, 74) // when append more than len -> algorithm will auto allocate
	// length a will increase to 11 while lenght x still unchange and cap a increase to 20 by algorithm
	fmt.Println("x: \t", x)
	fmt.Println("a: \t", len(a), cap(a), a)
	fmt.Println("b: \t", len(b), cap(b), b)

}
