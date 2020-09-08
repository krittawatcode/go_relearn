package main

import (
	"fmt"
	"time"
)

func stopWatch(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("function %s : took time %v\n", name, time.Now().Sub(start).Seconds())
	}
}

func loadingImage() {

	defer stopWatch("loadingImage")()
	time.Sleep(3 * time.Millisecond)
	fmt.Println("loadingImage : done")

}

func main() {
	// defer
	loadingImage()
}
