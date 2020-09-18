package main

import "fmt"

func sender(ch chan int, done chan string) {
	for i := 0; i <= 5; i++ {
		fmt.Println("Sending value : ", i)
		ch <- i
	}
	close(ch)
	fmt.Println("Sender is about to complete.")
	done <- "Done from Sender"
	fmt.Println("Sender done")
}

func receiver(ch chan int, done chan string) {
	for v := range ch {
		fmt.Println("\tReceive value : ", v)
	}
	done <- "Done from receiver"
}

func main() {
	ch := make(chan int)
	done := make(chan string)
	go sender(ch, done)
	go receiver(ch, done)
	doneStatus := <-done
	fmt.Println("Done status : ", doneStatus)
	doneStatus = <-done
	fmt.Println("Done status : ", doneStatus)
	fmt.Println("Main exit")
}
