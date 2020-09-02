package main

import (
	"fmt"

	gh "github.com/krittawatcode/weight"
)

func main() {
	k := gh.KG(1)
	fmt.Println(k)
	fmt.Println(gh.KgToLb(k))
}
