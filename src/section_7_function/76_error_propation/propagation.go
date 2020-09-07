package main

import (
	"errors"
	"fmt"
	"log"
)

const dbReady = true
const balance = 200

func getBalance() (int, error) {
	if !dbReady {
		return 0, errors.New("getBalance: database is down")
	}
	return balance, nil
}

func withdraw(amount int) (int, error) {

	balance, err := getBalance()

	if err != nil {
		return 0, fmt.Errorf("withdraw: %v", err)
	}

	if amount > balance {
		return 0, errors.New("withdraw: insufficient fund") // nil can't cast to 0
	}
	return amount, nil
}

func main() {
	log.SetFlags(0)
	amount, err := withdraw(400)
	if err != nil {
		log.Fatalf("main: %v", err)
	}
	fmt.Println("Please collect your money : ", amount)
}
