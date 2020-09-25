package main

import "fmt"

func main() {
	result := isPalindrome(12421)
	fmt.Println(result)
	testarr := []int{1, 2, 4, 2, 1}
	lenght := len(testarr) / 2
	fmt.Println(lenght)
}

// func isPalindrome(x int) bool {
// 	if x < 0 {
// 		return false
// 	}
// 	magic := []int8{}
// 	for x != 0 {
// 		magic = append(magic, int8(x%10))
// 		x /= 10
// 	}
// 	for i := 0; i < len(magic)/2; i++ {

// 		if magic[i] != magic[len(magic)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	a := make([]int, 0, 100)
	i := 0
	for x != 0 {
		a = append(a, x%10)
		x /= 10
		i++
	}
	for j := 0; j < i/2; j++ {
		if a[j] != a[i-1-j] {
			return false
		}
	}
	return true
}
