package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	// name        string
	// age         int
	person      Person // use this instead of name , age
	designation string
}

func main() {
	filicity := Employee{
		person:      Person{"Filicity", 23},
		designation: "Programmer",
	}

	// fmt.Println(filicity)
	fmt.Printf("%+v\n", filicity)
	fmt.Println(filicity.person.name)

}
