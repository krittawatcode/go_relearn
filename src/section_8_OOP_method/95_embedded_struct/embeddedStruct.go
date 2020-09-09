package main

import "fmt"

type Person struct {
	Name    string
	Surname string
}

func (p *Person) Fullname() string {
	return p.Name + " " + p.Surname
}

type Employee struct {
	Person
	Id     string
	Office string
}

func (e *Employee) Detail() string {
	return "ID: " + e.Id + ". Office: " + e.Office + ". Fullname: " + e.Fullname() + "."
}

func (e *Employee) IsSameOffice(other *Employee) bool {
	return e.Office == other.Office
}

type Programmer struct {
	Employee
	language []string
}

func (p *Programmer) Detail() string {
	return "Programmer : " + p.Employee.Detail()
}

type Tester struct {
	Employee
	Tools []string
}

func (t *Tester) Detail() string {
	return "Tester : " + t.Employee.Detail()
}

func main() {

	prach := Person{
		Name:    "Krittawat",
		Surname: "LOL go go go",
	}

	empPrach := Employee{
		Person: prach,
		Id:     "123",
		Office: "Thailand",
	}

	progPrach := Programmer{
		Employee: empPrach,
		language: []string{"golang", "java", "Java Script", "Type Script", "C#"},
	}

	oliver := Person{
		Name:    "ESher",
		Surname: "Queen",
	}

	empOliver := Employee{
		Person: oliver,
		Id:     "456",
		Office: "Thailand",
	}

	testerOliver := Tester{
		Employee: empOliver,
		Tools:    []string{"Robot"},
	}

	fmt.Printf("%+v\n", progPrach) // %+v = include field
	fmt.Printf("%+v\n", testerOliver)

	fmt.Println(empPrach.IsSameOffice(&empOliver))
	fmt.Println(progPrach.IsSameOffice(&testerOliver.Employee)) // can't use tester because this method use in Employee, so we need to .Employee to use this method
	fmt.Println(progPrach.Fullname())
	fmt.Println(progPrach.Detail())

	fmt.Println(progPrach.Employee.Detail())
}
