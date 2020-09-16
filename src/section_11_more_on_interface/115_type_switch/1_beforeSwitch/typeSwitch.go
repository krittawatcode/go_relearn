package main

import "fmt"

func testValue(x interface{}) {
	if v, ok := x.(string); ok {
		fmt.Println("string value: ", v)
		return
	}
	if v, ok := x.(int); ok {
		fmt.Println("int value: ", v)
		return
	}
	if v, ok := x.(bool); ok {
		fmt.Println("bool value: ", v)
		return
	}
	if v, ok := x.(float32); ok {
		fmt.Println("float32 value: ", v)
		return
	}
	if v, ok := x.(float64); ok {
		fmt.Println("float64 value: ", v)
		return
	}
	fmt.Println("No match any")
}

func main() {
	testValue("sdf")           // string
	testValue(264)             // int
	testValue(true)            // boolean
	testValue(float32(264.23)) // float32
	testValue(264.23)          // float64
}
