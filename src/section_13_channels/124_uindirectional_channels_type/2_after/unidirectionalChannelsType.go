package main

import "fmt"

func sender(ch chan<- int) {
	ch <- 88
	// f(ch)
}

func f(ch chan int) {
	ch <- 99
}

func receiver(ch <-chan int, done chan bool) {
	fmt.Println("Receive ", <-ch)
	done <- true
}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go sender(ch)
	go receiver(ch, done)
	<-done
	fmt.Println("Done")
}
