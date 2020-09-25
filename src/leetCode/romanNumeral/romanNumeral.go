package main

import "fmt"

/*
	https://leetcode.com/problems/roman-to-integer/
*/

func main() {
	x := romanToInt("MCMXCIV")
	// x := romanToInt("III")
	// x := romanToInt("IV")
	fmt.Println(x)

}

var m = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var result int

	prev := s[0]
	value := m[prev]
	// fmt.Println(len(s))
	for i := 1; i < len(s); i++ {
		c := s[i]
		v := m[c]

		if v > m[prev] {
			prev = c
			result += v - value
			value = 0

			continue
		}

		prev = c
		result += value
		value = v
	}

	return result + value
}

// func romanToInt(s string) int {
// 	// s = IV
// 	var result int

// 	start := s[0]   // I
// 	val := m[start] // 1

// 	for i := 1; i < len(s); i++ {
// 		second := s[i]          // V
// 		currentVal := m[second] // 5
// 		if currentVal > val {
// 			result += currentVal - val // result = 5 - 1 = 4
// 			continue
// 		} else {
// 			result += currentVal
// 		}
// 	}

// 	return result
// }
