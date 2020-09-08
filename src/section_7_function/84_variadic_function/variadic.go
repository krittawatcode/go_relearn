package main

import "fmt"

func printEachLine(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	// printEachLine("abc", "def", 123)

	x := []interface{}{"abc", "def", 234, "ghi"}
	// printEachLine(x)
	printEachLine(x...)
}
