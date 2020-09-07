package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

const dbReady = false
const balance = 200
const numberToSuccess = 2000

func connectDb(nTry int) error {
	if nTry == numberToSuccess {
		return nil
	}
	return errors.New("busy")
}

func waitForDatabase() error {

	timeout := 3 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		err := connectDb(tries)
		if err == nil {
			// log.Panicln(err)
			return nil
		}
		log.Printf("database is not responding (%v); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}

	// err := connectDb(1)
	// if err != nil {
	// 	log.Println(err)
	// }
	// connectDb(2)
	// if err != nil {
	// 	log.Println(err)
	// }
	// return nil
	return fmt.Errorf("waitForDatabase: timeout %v", timeout)
}

func getBalance() (int, error) {
	if !dbReady {
		// log.Println("retrying....")
		if err := waitForDatabase(); err != nil {
			return 0, fmt.Errorf("getBalance: %v", err)
		}
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
	amount, err := withdraw(200)
	if err != nil {
		log.Fatalf("main: %v", err)
	}
	fmt.Println("Please collect your money : ", amount)
}
