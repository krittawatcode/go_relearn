package main

import (
	"fmt"
	"os"
)

type Notebook struct {
	content []byte
}

func (nb *Notebook) String() string {
	return string(nb.content)
}

func (nb *Notebook) Write(p []byte) (n int, err error) {
	nb.content = append(nb.content, p...)
	return len(p), nil
}

func main() {
	// fmt.Println("Hello")
	fmt.Fprintln(os.Stdout, "Hello World") // same as fmt.Println("Hello")
	nb := Notebook{}
	fmt.Println(nb)
	fmt.Fprintln(&nb, "hello world")
	fmt.Println(&nb)
}
