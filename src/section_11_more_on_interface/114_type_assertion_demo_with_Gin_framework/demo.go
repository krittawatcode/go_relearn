package main

import (
	"fmt"
)

func main() {
	Value("sdf")
	Value(45)
}

func Value(key interface{}) {
	if keyAsString, ok := key.(string); ok {
		fmt.Println("assert ok : value is ", keyAsString)
	} else {
		fmt.Println("assert not ok")
	}
}
