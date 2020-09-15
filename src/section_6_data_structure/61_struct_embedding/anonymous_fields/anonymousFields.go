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
type CEO struct {
	Employee
	commandpower string
}

func main() {
	filicity := Employee{
		Person{"Filicity", 23},
		"Programmer",
	}

	fmt.Printf("%+v\n", filicity)
	fmt.Println(filicity.name)

	prach := CEO{

		Employee{
			Person{
				"Prach", 26,
			},
			"Programmer",
		},
		"demacia",
	}
	fmt.Println(prach)
}
