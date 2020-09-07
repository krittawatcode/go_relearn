package main

import "fmt"

func createF(list []int) []func() {
	result := []func(){}

	for i, v := range list {
		capI := i
		capV := v
		result = append(result, func() {
			fmt.Println(capI, capV)
		})
	}

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
