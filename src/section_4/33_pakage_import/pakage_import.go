package main

import (
	"fmt"
	"weight"
)

func main() {
	k := weight.KG(1)
	fmt.Println(k)
	fmt.Println(weight.KgToLb(k))
}
