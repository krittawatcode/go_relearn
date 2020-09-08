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

	david := Person{
		Name:    "Devid",
		Surname: "GG",
	}

	empDavid := Employee{
		Person: david,
		Id:     "123",
		Office: "Thailand",
	}

	progDavid := Programmer{
		Employee: empDavid,
		language: []string{"golang", "java"},
	}

	oliver := Person{
		Name:    "Oliver",
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

	fmt.Printf("%+v\n", progDavid) // %+v = include field
	fmt.Printf("%+v\n", testerOliver)

	davidDetail := progDavid.Detail              // set but don't use
	fmt.Println("davidDetail : ", davidDetail()) // use at this point

	// fmt.Println((*Employee).IsSameOffice(&empDavid, &empOliver)) // pass receiver and parameter

	IsSameOffice := (*Employee).IsSameOffice
	fmt.Println(IsSameOffice(&empDavid, &empOliver))

}
