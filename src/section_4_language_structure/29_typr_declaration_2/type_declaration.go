package main

import "fmt"

type KG float64
type LB float64

func (lb LB) toKg() KG {
	return KG(lb / 2.2046226218)
}

func main() {

	k := KG(3)
	lb := LB(3)

	fmt.Println(k > lb.toKg())

}
