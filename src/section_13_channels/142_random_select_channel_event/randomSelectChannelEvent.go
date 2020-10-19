package main

import "fmt"

func running(name string, ch <-chan int) {
	for {
		fmt.Println(name, <-ch)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go running("Ch1", ch1)
	go running("Ch2", ch2)
	go running("Ch3", ch3)
	go running("Ch4", ch4)

	for i := 0; i < 10; i++ {
		select {
		case ch1 <- i:
		case ch2 <- i:
		case ch3 <- i:
		case ch4 <- i:
		}
	}
}
