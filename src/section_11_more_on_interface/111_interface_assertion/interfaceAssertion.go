package main

import (
	"fmt"
	"io"
	"os"
)

/*
	https://tour.golang.org/methods/15

	Type assertions
	A type assertion provides access to an interface value's underlying concrete value.

	t := i.(T)
	This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

	If i does not hold a T, the statement will trigger a panic.

	To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

	t, ok := i.(T)
	If i holds a T, then t will be the underlying value and ok will be true.

	If not, ok will be false and t will be the zero value of type T, and no panic occurs.

	Note the similarity between this syntax and that of reading from a map.

*/

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
