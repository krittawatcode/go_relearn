package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// training() -- we dont know how long it will finish training -> we can't use time.sleep
	done := make(chan bool)
	go training(done)
	trainingStatus := <-done // send value from channel to var
	fmt.Println("Done status : ", trainingStatus)
}

func training(ch chan bool) {
	rand.Seed(int64(time.Now().Nanosecond()))
	x := rand.Intn(1000)
	fmt.Println("training for : ", x, " millisecond")
	time.Sleep(time.Duration(x) * time.Millisecond)
	ch <- true // send value back to channel
}
