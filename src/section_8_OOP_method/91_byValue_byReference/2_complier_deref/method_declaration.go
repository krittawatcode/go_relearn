package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Length() float64 {
	return math.Hypot((*p).X, (*p).Y)
}

func (p *Point) MoveYTo(newY float64) {
	MoveYTo(p, newY)
}

//  ----- by value -----
func (p Point) MoveYToByValue(newY float64) {
	MoveYTo(&p, newY)
}

// normal function
func MoveYTo(p *Point, newY float64) {
	(*p).Y = newY
}

func main() {
	p := Point{3, 4}
	(&p).MoveYToByValue(31) // complier will *(&p)
	(&p).MoveYTo(9)
	fmt.Println(p)
}
