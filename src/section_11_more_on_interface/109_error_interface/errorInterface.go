package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	msg string
}

func (m MyError) Error() string {
	return m.msg
}

func main() {
	// ioutil.ReadFile()
	var err error
	err = errors.New("223")
	fmt.Printf("%T , %#v\n", err, err)                  // *errors.errorString , &errors.errorString{s:"223"}
	fmt.Println(errors.New("223") == errors.New("223")) // false

	err = MyError{"this is an error"}
	fmt.Printf("%T , %#v\n", err, err)
	// fmt.Errorf("", "")
}
