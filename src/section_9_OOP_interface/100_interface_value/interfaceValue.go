package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer               // zero
	fmt.Printf("%T , %v\n", w, w) // nil nil

	w = os.Stdout                // concrete type
	fmt.Printf("%T, %v\n", w, w) // *os.File, &{0xc000096000}

	w = &bytes.Buffer{}          // 0 value
	fmt.Printf("%T, %v\n", w, w) // 0 value
	w.Write([]byte("hello"))     // change 0 value to hello
	fmt.Printf("%T, %v\n", w, w) // hello

	w = nil
	fmt.Printf("%T, %v\n", w, w) // nil, nil

	var byteBuff *bytes.Buffer
	// printBuff(byteBuff)
	fmt.Println(byteBuff == nil)
	fmt.Println(w == nil)
	w = byteBuff          // assign nil = nil
	fmt.Println(w == nil) // when you check nil it will check both type and value, So ->
	// false bcz it will change to *bytes.Buffer, <nil> (from <nil>, <nil>)
	fmt.Printf("%T, %v\n", w, w) // check nil of interface -> when type and value 'nil' -> true

	/*
		func printBuff(w io.Writer){
			if w != nil {
				fmt.Println(w.Write([]byte("hello")))
			}
		}
	*/

}
