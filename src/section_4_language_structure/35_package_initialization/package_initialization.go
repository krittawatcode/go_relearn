package main

import (
	"fmt"
	"math/rand"
	"prime"
)

func main() {
	fmt.Println(prime.IsPrime(2))
	fmt.Println(prime.IsPrime(30))
	fmt.Println(prime.IsPrime(390341))

	for i := 0; i <= 100; i++ {
		x := rand.Intn(1000000)
		fmt.Printf("%d , %t\n", x, prime.IsPrime(x))
	}
}
