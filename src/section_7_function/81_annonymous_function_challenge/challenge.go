package main

import "fmt"

func createF(list []int) []func() {
	result := []func(){}

	for _, v := range list {
		result = append(result, func() {
			fmt.Println(v)
		})
	}

	// result = append(result, func() {
	// 	fmt.Println(list[0])
	// })

	// result = append(result, func() {
	// 	fmt.Println(list[1])
	// })

	return result

}

func main() {
	fList := []func(){}
	fList = createF([]int{108, 12, 30, 4, 5})
	fmt.Println("-----------")
	for _, v := range fList {
		v()
	}
}
