package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type byName []*Person

type customSort struct {
	Persons []*Person
	less    func(i, j *Person) bool
}

func (p customSort) Len() int {
	return len(p.Persons)
}

func (p customSort) Less(i, j int) bool {
	return p.less(p.Persons[i], p.Persons[j])
}

func (p customSort) Swap(i, j int) {
	p.Persons[i], p.Persons[j] = p.Persons[j], p.Persons[i]
}

// type Interface interface {
// Len is the number of elements in the collection.
func (p byName) Len() int {
	return len(p)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p byName) Less(i, j int) bool {
	return p[i].name < p[j].name
}

// Swap swaps the elements with indexes i and j.
func (p byName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// }

func main() {

	fmt.Println("Sort")
	p := []*Person{
		{"A", 22},
		{"B", 21},
		{"B", 24},
		{"C", 25},
		{"C", 20},
		{"A", 21},
		{"A", 22},
		{"B", 22},
	}

	// fmt.Println(p)
	printPerson(p)
	sort.Sort(customSort{Persons: p, less: func(i, j *Person) bool {
		if i.name != j.name {
			return i.name < j.name
		}
		if i.age != j.age {
			return i.age < j.age
		}
		return false
	}})
	printPerson(p)
}

func printPerson(p []*Person) {
	fmt.Println("--------------------")
	for _, v := range p { // for index, value := range
		fmt.Println(*v)
	}
}
