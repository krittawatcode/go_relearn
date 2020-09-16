package main

import "fmt"

func testValue(x interface{}) {

	switch v := x.(type) {
	case string:
		fmt.Println("string value: ", v)
	case int:
		fmt.Println("int value: ", v)
	case bool:
		fmt.Println("bool value: ", v)
	case float32, float64:
		fmt.Println("float value: ", v)
	default:
		fmt.Println("No match any")
	}

	/* // or use this if you don't want to use value
	switch x.(type) {
	case string:
		fmt.Println("string value: ")
	case int:
		fmt.Println("int value: ")
	case bool:
		fmt.Println("bool value: ")
	case float32, float64:
		fmt.Println("float value: ")
	default:
		fmt.Println("No match any")
	}
	*/
}

func main() {
	testValue("sdf")           // string
	testValue(264)             // int
	testValue(true)            // boolean
	testValue(float32(264.23)) // float32
	testValue(264.23)          // float64
}
