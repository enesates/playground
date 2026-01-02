package main

import (
	"fmt"
)

type InsufficientFunds struct{}

func (fund *InsufficientFunds) Error() string {
	return "InsufficientFunds"
}

type CustomError struct{}

func (fund *CustomError) Error() string {
	return "CustomError"
}

type BankAccount struct {
	balance float64
}

func (ba *BankAccount) Withdraw(amount float64) error {
	if amount > ba.balance {
		return &InsufficientFunds{}
	}

	ba.balance -= amount

	return nil
}

func main() {
	ba := BankAccount{10}

	if err := ba.Withdraw(5); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Balance:", ba.balance)
}
