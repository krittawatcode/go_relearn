package main

import (
	"errors"
	"fmt"
)

func withdraw(amount int) (int, error) {
	if amount > 100 {
		// os.Exit(1) -> next session
		return 0, errors.New("insufficient fund") // nil can't cast to 0
	}
	return amount, nil
}

func main() {
	amount, err := withdraw(200)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Please collect your money : ", amount)
}
