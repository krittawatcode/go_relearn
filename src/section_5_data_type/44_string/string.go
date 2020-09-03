package main

import "fmt"

func main() {
	eng := "Hello"
	th := "สวัสดี"
	fmt.Println(eng, len(eng), eng[0]) // 5
	fmt.Println(th, len(th), th[0])    // 18

	fmt.Println(eng, len(eng), string(eng[0])) // 5
	fmt.Println(th, len(th), string(th[0]))    // 18
}
