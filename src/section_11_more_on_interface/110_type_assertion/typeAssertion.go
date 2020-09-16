package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// x.(T)
	// key.(string)

	/*
		interface.(type/interface) <- asserted type
	*/

	var w io.Writer        // Write(p []byte) (n int, err error)
	w = os.Stdout          // will return pointer of file
	result := w.(*os.File) // *os.File & os.Stdout both imp io.Writer so we can assert it
	// result := w.(*bytes.Buffer) // panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer
	fmt.Printf("%T , %#v\n", w, w) // will print type of dynamic value of w
	fmt.Printf("%T , %#v\n", result, result)
	// value will point back to os.Stdout value

}
