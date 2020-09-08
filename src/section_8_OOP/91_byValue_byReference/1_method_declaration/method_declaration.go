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
	p.MoveYToByValue(31)
	p.MoveYTo(9)            // complier will automatic ref of p to method -> can use p.MoveYTo()
	(&p).MoveYToByValue(31) // complier will autometic change &p to p -> then pass only value to method
	fmt.Println(p)
}
