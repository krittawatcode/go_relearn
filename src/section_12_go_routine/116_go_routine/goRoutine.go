package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("1")
	go fmt.Println("2")
	go fmt.Println("3")
	// go panic("test")
	go fmt.Println("4")
	time.Sleep(1000)
	/*
		main routine = do all job then stop
		but when use 'go' its will count as succsee even task don't finish
	*/
}
