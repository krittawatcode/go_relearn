package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	s := a[1:3]
	S := a[:3]                     // start at index 0 exclude index 3
	fmt.Println(len(s), cap(s), s) // 2 3 [2 3]
	fmt.Println(len(S), cap(S), S) // 3 4 [1 2 3]

	ss := make([]int, 5, 10) // make([]int, len, cap) or make([]int, len)
	// ss[0] = 99
	// ss[6] = 99
	fmt.Println(len(ss), cap(ss), ss)

	// if out of memory -> panic: runtime error: index out of range
}
