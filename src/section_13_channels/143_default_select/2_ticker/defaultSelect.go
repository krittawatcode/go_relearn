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
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go running("Ch1", ch1)
	go running("Ch2", ch2)
	go running("Ch3", ch3)
	go running("Ch4", ch4)

	ticker1 := time.NewTicker(5 * time.Millisecond)
	ticker2 := time.NewTicker(6 * time.Millisecond)
	ticker3 := time.NewTicker(3 * time.Millisecond)
	ticker4 := time.NewTicker(4 * time.Millisecond)

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
		default:
			fmt.Println("Default")
		}
		time.Sleep(1 * time.Millisecond)
	}
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Done")
}
