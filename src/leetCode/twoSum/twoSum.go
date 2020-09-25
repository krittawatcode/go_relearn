package main

import "fmt"

func main() {
	data := []int{2, 7, 11, 15}
	result := twoSum(data, 18)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	// target = 9
	m := make(map[int]int)
	for i, v := range nums {
		// i = 0 , v = 2
		// i = 1 , v = 7
		if j, ok := m[target-v]; ok {
			//
			return []int{j, i}
		} else {
			// key 2 value 0
			m[v] = i
		}
	}
	return []int{-1, -1}
}
