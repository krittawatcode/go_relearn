package main

import (
	"fmt"
	"time"
)

func running(name string, ch <-chan int) {
	for {
		fmt.Println(name, <-ch)
	}
}

func main() {

	after2Second := time.After(2 * time.Second)

	/*
		ch3 will run first because ticker3 will finish first
	*/
	for i := 0; i < 10; i++ {
		select {
		case v := <-after2Second:
			fmt.Println("got value after 2 sec", v.Second())
		default:
			fmt.Println("Default")
		}
		time.Sleep(500 * time.Millisecond)
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Done")
}
