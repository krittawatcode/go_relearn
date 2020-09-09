package main

import (
	"flag"
	"fmt"
	"strconv"
)

var romanMap = map[string]int{
	"i":   1,
	"ii":  2,
	"iii": 3,
	"iv":  4,
	"v":   5,
}

type RomanAge struct {
	age string
}

func (r RomanAge) String() string {
	return strconv.Itoa(romanMap[r.age])
}

func (r *RomanAge) Set(value string) error {
	r.age = value
	return nil
}

func flagRomanAge() *RomanAge {
	romanAge := RomanAge{}
	flag.Var(&romanAge, "age", "help message for flagname")
	return &romanAge
}

var name = flag.String("name", "anonymous", "your name") // input => flag name, default value, help message // return => address of that flag name

var age = flagRomanAge()

func main() {

	/* Important!!!!
	run this code with command line
	go run '.\src\section_9_OOP_interface\99_flag_interface\flagInterface.go' -name Filicity
	*/

	flag.Parse()
	// fmt.Println(name)  // return address
	// fmt.Println(*name) // return value

	/* run without flag name */
	// fmt.Println(*name) // will return anonymous

	/* last part */
	/*
		run this code with command line
		go run '.\src\section_9_OOP_interface\99_flag_interface\flagInterface.go' -name Filicity -age iv
	*/
	fmt.Println(*age)
}
