package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const N = 1000
	add := 0
	sub := 0
	mu := sync.Mutex{} // mutual exclusion
	for i := 0; i < N; i++ {
		add++
		go func() {
			defer func() {
				mu.Lock()
				sub-- // critical session
				mu.Unlock()
			}()
		}()
	}

	for {
		h, m, s := time.Now().Clock()
		// fmt.Println(h, m, s)
		fmt.Printf("%d-%d-%d : add %d , sub %d\r", h, m, s, add, sub)
		if add == N && sub == -N {
			fmt.Println("success")
			break
		}
	}
	fmt.Println("Main done : ", add, sub)
}
