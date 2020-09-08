package main

import (
	"fmt"
	"os"
)

var cygwinPath = os.Getenv("CYGWIN")

func init() {
	defer func() {
		fmt.Println("clean up in init") // run last but before panic
	}()
	fmt.Println("init cygwin path : ", cygwinPath)
	if cygwinPath == "" {
		panic("cannot set cygwin path. Make sure you have one.")
	}
}

func main() {
	// utility for Cygwing
	// path cygwing?
	fmt.Println("Hello to Cygwing utilitys")

}
