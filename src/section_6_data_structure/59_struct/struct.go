package main

import "fmt"

func main() {

	/* ---- I ---- */
	// x := struct {
	// 	name string
	// 	age  int
	// }{"filicity", 23}

	/* ---- II ---- */
	x := struct {
		name string
		age  int
	}{age: 23, name: "filicity"}

	fmt.Println(x, x.name, x.age)
}
