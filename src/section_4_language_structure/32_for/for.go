package main

import (
	"fmt"
	"time"
)

func main() {

	// for init; condition; post {
	// 	body
	// }

	for i := 1; i < 5; i++ {
		fmt.Println(time.Now().Clock())
	}

	// with break
	for i := 1; i < 5; i++ {
		fmt.Println(i)
		if i == 3 {
			break
		}
	}

	// with continune
	for i := 1; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("skip", i)
			continue
		}
	}
}
