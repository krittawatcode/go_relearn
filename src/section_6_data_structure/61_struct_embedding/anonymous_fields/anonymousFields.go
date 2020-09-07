package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	Person
	designation string
}

func main() {
	filicity := Employee{
		Person{"Filicity", 23},
		"Programmer",
	}

	fmt.Printf("%+v\n", filicity)
	fmt.Println(filicity.name)

}
