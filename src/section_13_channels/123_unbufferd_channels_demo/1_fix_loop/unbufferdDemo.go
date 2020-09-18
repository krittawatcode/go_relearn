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

	/*
		fmt.Println("\tReceive value : ", <-ch) // 0
		fmt.Println("\tReceive value : ", <-ch) // 1
		fmt.Println("\tReceive value : ", <-ch) // 2
		fmt.Println("\tReceive value : ", <-ch) // 3
		fmt.Println("\tReceive value : ", <-ch) // 4
		fmt.Println("\tReceive value : ", <-ch) // 5
		// fmt.Println("\tReceive value : ", <-ch) // 6 will make deadlock if no sender
		fmt.Println("\tReceive value : ", <-ch) // 0
		fmt.Println("\tReceive value : ", <-ch) // 0
		fmt.Println("\tReceive value : ", <-ch) // 0
		done <- "Done from receiver"
	*/

	for i := 0; i <= 9; i++ {
		fmt.Println("\tReceive value : ", <-ch)
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
