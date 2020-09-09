package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Length(p *Point) float64 {
	return math.Hypot((*p).X, (*p).Y)
}

func MoveXTo(p *Point, newX float64) {
	(*p).X = newX
}

func MoveYTo(p *Point, newY float64) {
	(*p).Y = newY
}

func main() {
	p := Point{3, 4}
	fmt.Println(p)
	fmt.Println(Length(&p)) // 5
	MoveXTo(&p, 6)
	fmt.Println(Length(&p)) // 7.2
}
