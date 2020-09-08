package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Length() float64 {
	// return math.Hypot((*p).X, (*p).Y)
	return math.Hypot(p.X, p.Y)
}

func (p *Point) MoveYTo(newY float64) {
	MoveYTo(p, newY)
}

func MoveYTo(p *Point, newY float64) {
	p.Y = newY
}

func main() {

	/*
		when we create method to struct, method receiver will have same type in every method
		-> if pointer then every method will use pointer receiver every method even we don't want to change value in strut
	*/

	p := Point{3, 4}

	(&p).MoveYTo(31)
	fmt.Println(p)
}
