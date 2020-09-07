package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// count
	x := "ทดสอบ"
	finder := "สอ"
	fmt.Println(bytes.Count([]byte(x), []byte{})) // empty count then return 1 +
	fmt.Println(bytes.Count([]byte(x), []byte(finder)))

	// byte buffer
	buff := bytes.Buffer{}
	buff.WriteRune('h')
	buff.WriteRune('i')
	fmt.Println(buff.String())

	// builder
	buf := strings.Builder{}
	buf.WriteRune('h')
	buf.WriteRune('i')
	fmt.Println(buf.String())

	// convert
	atoi, _ := strconv.Atoi("123") // string to int
	itoa := strconv.Itoa(123)
	fmt.Println(atoi, reflect.TypeOf(atoi))
	fmt.Println(itoa, reflect.TypeOf(itoa))

	// convert boolean
	fmt.Println(strconv.ParseBool("True"))
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("False"))
	fmt.Println(strconv.ParseBool("FalSe"))
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseBool("0"))

	// digit
	fmt.Println("unicode.IsDigit('1') : ", unicode.IsDigit('1'))
}
