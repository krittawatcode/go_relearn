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

	prev := s[0]     // M // C
	value := m[prev] // 1000
	// fmt.Println(len(s))
	for i := 1; i < len(s); i++ { // i = 1 // i = 2
		c := s[i] // C // M
		v := m[c] // 100 // 1000

		if v > m[prev] { // 100 1000 // 1000 100
			prev = c            // M
			result += v - value // 1000 + 1000 - 100
			value = 0

			continue
		}

		prev = c        // C
		result += value // 1000
		value = v       // 100
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
