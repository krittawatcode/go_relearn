package main

import "fmt"

func main() {
	x := 1
	// switch case will do when it finish case
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// switch with fallthrough
	y := 200 // -20
	switch { // omit condition true == switch true
	case 1 <= y && y <= 10:
		fmt.Println("deci")
		fallthrough
	case 100 <= y && y <= 1000:
		fmt.Println("centi")
		fallthrough
	case 1000 <= y:
		fmt.Println("meter")
	default:
		fmt.Println("default")
	}

	// switch with break
	z := "f" // d z
	switch z {
	case "a":
		fmt.Println("A")
	case "b":
		fmt.Println("B")
	case "c":
		fmt.Println("C")
	case "d", "e", "f", "g":
		fmt.Println("D E F G")
		if z == "f" {
			break
		}
		fmt.Println("your lucky")
	default:
		fmt.Println("default")
	case "z":
	}
}
