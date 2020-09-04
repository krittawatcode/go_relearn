package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	x := person{age: 23, name: "filicity"}
	y := person{age: 23}
	z := person{age: 23, name: "filicity"}

	fmt.Println(x, x.name, x.age) // print only value
	fmt.Printf("%+v", x)          // +v will print key(field name) value
	fmt.Println(y)
	fmt.Println(x == z) // compare value not address
	fmt.Println(&x, &z) // ??
}
