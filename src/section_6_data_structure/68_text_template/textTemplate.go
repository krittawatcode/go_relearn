package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func upperCase(s string) string {
	return strings.ToUpper(s)
}

func main() {
	filicity := Person{"Filicity", 23}
	oliver := Person{"Oliver", 25}

	people := []Person{filicity, oliver}

	const greetPerson = `Hi I am {{.Name | upperCase }}. {{.Age}} years old{{"\n"}}`
	const greetPeople = `{{range .}}Hi I am {{.Name}}. {{.Age}} years old{{"\n"}}{{end}}`

	/* ---- greetTemplate case I ---- */
	// greetTemplate := template.Must(template.New("greetingFromPerson").Funcs(template.FuncMap{"upperCase": upperCase}).Parse(greetPerson))

	/* ---- greetTemplate case II ---- */
	maps := make(template.FuncMap) // template.FuncMap{"upperCase": upperCase}
	maps["upperCase"] = upperCase
	greetTemplate := template.Must(template.New("greetingFromPerson").Funcs(maps).Parse(greetPerson))

	greetPeopleTemplate := template.Must(template.New("greetingFromPerson").Parse(greetPeople))

	fmt.Println(greetTemplate.Name())
	greetTemplate.Execute(os.Stdout, filicity)
	greetTemplate.Execute(os.Stdout, oliver)
	fmt.Println("==================================")
	greetPeopleTemplate.Execute(os.Stdout, people)
}
