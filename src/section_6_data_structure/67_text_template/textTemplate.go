package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	filicity := Person{"Filicity", 23}
	oliver := Person{"Oliver", 23}

	const greetPerson = `Hi I am {{.Name}}. {{.Age}} years old{{"\n"}}`
	// fmt.Printf("Hi I am %s. %d years old\n", filicity.Name, filicity.Age)
	// fmt.Printf("Hi I am %s. %d years old\n", oliver.Name, oliver.Age)

	/* ----- first way ----- */
	// greetTemplate, err := template.New("greetingFromPerson").Parse(greetPerson)
	// if err != nil {
	// 	return
	// }

	greetTemplate := template.Must(template.New("greetingFromPerson").Parse(greetPerson))

	fmt.Println(greetTemplate.Name())
	greetTemplate.Execute(os.Stdout, filicity)
	greetTemplate.Execute(os.Stdout, oliver)

}
