package main

import "fmt"

type Square struct {
	Length int
}

func (s Square) Area() int {
	return s.Length * s.Length
}

func (s Square) Perimeter() int {
	return s.Length * 4
}

//////////////////////////////////

type BankAccount struct {
	Balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount int) {
	b.Balance -= amount
}

func main() {
	square := Square{Length: 3}
	fmt.Printf("Area of square (%d): %d\n", square.Length, square.Area())
	fmt.Printf("Perimeter of square (%d): %d\n", square.Length, square.Perimeter())

	//////////////////////////////////

	bankAccount := BankAccount{Balance: 100}
	fmt.Printf("Bank account balance: %d\n", bankAccount.Balance)

	depositAmount := 200
	bankAccount.Deposit(depositAmount)
	fmt.Printf("Bank account balance after deposit (%d): %d\n", depositAmount, bankAccount.Balance)

	withdrawAmount := 50
	bankAccount.Withdraw(withdrawAmount)
	fmt.Printf("Bank account balance after withdraw (%d): %d\n", withdrawAmount, bankAccount.Balance)
}
