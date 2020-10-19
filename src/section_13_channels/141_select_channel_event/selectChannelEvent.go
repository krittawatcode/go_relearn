package main

import (
	"fmt"
	"math/rand"
	"time"
)

func running(track chan<- struct{}) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	track <- struct{}{}
}

func main() {
	track1 := make(chan struct{})
	track2 := make(chan struct{})
	track3 := make(chan struct{})

	go running(track1)
	go running(track2)
	go running(track3)

	select {
	case <-track1:
		fmt.Println("Winner is Hourse 1")
	case <-track2:
		fmt.Println("Winner is Hourse 2")
	case <-track3:
		fmt.Println("Winner is Hourse 3")
	}

	// for {
	// 	select {
	// 	case <-track1:
	// 		fmt.Println("Winner is Hourse 1")
	// 	case <-track2:
	// 		fmt.Println("Winner is Hourse 2")
	// 	case <-track3:
	// 		fmt.Println("Winner is Hourse 3")
	// 	}
	// }

	fmt.Println("Done")
}
