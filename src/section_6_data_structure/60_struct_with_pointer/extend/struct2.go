package main

import "fmt"

// struct (record)
type Person struct {
	name string
	age  int
}

func main() {

	x := Person{age: 24, name: "filicity"}
	// z := x // pass by value
	// z.name = "Prach"
	z := &x // pass by reference
	z.name = "Prach"
	fmt.Println(x, x.name, x.age)
	fmt.Println(z)

	/* ----- Part II : compare struct -----*/
	y := Person{age: 24, name: "Prach"}
	fmt.Println(x == y) // check value
	fmt.Printf("%+v\n", y)

	/* ----- Part III : new struct -----*/
	fmt.Println("--------------------")

	p := new(Person)         // p is pointer
	*p = Person{"Prach", 24} // assign value to pointer by de-reference -> *p = ....
	fmt.Println(p)           // p is still pointer
	fmt.Println(&p)          // show position in mem
	fmt.Println(*p)          // value of p

}
