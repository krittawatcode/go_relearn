package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func running(track chan<- struct{}) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(10*time.Second + time.Duration(rand.Intn(100))*time.Millisecond))
	track <- struct{}{}
}

func main() {
	defer func() {
		fmt.Println("Done")
	}()

	// Notify(c chan<- os.Signal, sig ...os.Signal)
	chInt := make(chan os.Signal, 1)
	signal.Notify(chInt, syscall.SIGINT, syscall.SIGHUP)

	track1 := make(chan struct{})
	track2 := make(chan struct{})
	track3 := make(chan struct{})
	abort := make(chan struct{})

	go running(track1)
	go running(track2)
	go running(track3)

	lapseChan := make(chan time.Time)

	go func() {
		lapseTricker := time.NewTicker(1 * time.Second)
		for {
			lapseChan <- (<-lapseTricker.C)
		}
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("Abort...")
		abort <- struct{}{}
	}()

	for done := false; !done; {
		select {
		case <-track1:
			fmt.Println("Winner is Horse 1")
			done = true
		case <-track2:
			fmt.Println("Winner is Horse 2")
			done = true
		case <-track3:
			fmt.Println("Winner is Horse 3")
			done = true
		case v := <-lapseChan:
			fmt.Println("Horses are running: ", v.Second())
		case <-abort:
			done = true
			fmt.Println("User pressing Enter")
		case v := <-chInt:
			done = true
			fmt.Printf("User interrupt : %v \n", v)
		}
	}

}
