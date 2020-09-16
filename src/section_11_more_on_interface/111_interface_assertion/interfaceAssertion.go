package main

import (
	"fmt"
	"io"
	"os"
)

type MyIo struct {
}

func (m MyIo) Read(p []byte) (n int, err error) {
	return len(p), nil
}
func (m MyIo) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func main() {
	var w io.Writer
	w = os.Stdout
	result := w.(MyIo) // panic: interface conversion: io.Writer is *os.File, not main.MyIo
	fmt.Printf("%T , %#v\n", w, w)
	fmt.Printf("%T , %#v\n", result, result)
	// w.Write
	// result.Read()
	// result.Write()

}
