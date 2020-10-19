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
	ticker1 := time.NewTicker(5 * time.Millisecond)
	ticker2 := time.NewTicker(6 * time.Millisecond)
	ticker3 := time.NewTicker(3 * time.Millisecond)
	ticker4 := time.NewTicker(4 * time.Millisecond)

	/*
		ch3 will run first because ticker3 will finish first
	*/
	for i := 0; i < 10; i++ {
		select {
		case v := <-ticker1.C:
			fmt.Println("ch1", v.Nanosecond())
		case v := <-ticker2.C:
			fmt.Println("ch2", v.Nanosecond())
		case v := <-ticker3.C:
			fmt.Println("ch3", v.Nanosecond())
		case v := <-ticker4.C:
			fmt.Println("ch4", v.Nanosecond())
			// default:
			// 	fmt.Println("Default")
		}
		time.Sleep(1 * time.Millisecond)
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Done")
}
