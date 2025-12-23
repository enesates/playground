package main

import "fmt"

type Payment interface {
	Process() string
}
type CreditCard struct {
	CreditCardNumber string
}

func (cc CreditCard) Process() string {
	return cc.CreditCardNumber
}

type Bitcoin struct {
	WalletNumber string
}

func (bt Bitcoin) Process() string {
	return bt.WalletNumber
}

type Bank struct {
	IBAN string
}

func (ba Bank) Process() string {
	return ba.IBAN
}

type Stock struct {
	Name string
}

func (st Stock) Process() string {
	return st.Name
}

func ProcessPayment(p Payment) {
	switch p.(type) {
	case CreditCard:
		fmt.Println("Credit Card payment has been processed:", p.Process())
	case Bitcoin:
		fmt.Println("Bitcoin payment has been processed:", p.Process())
	case Bank:
		fmt.Println("Bank payment has been processed:", p.Process())
	default:
		fmt.Println("Unsupported payment type")
	}
}

func main() {
	cc := CreditCard{CreditCardNumber: "123"}
	bc := Bitcoin{WalletNumber: "456"}
	ba := Bank{IBAN: "789"}
	st := Stock{Name: "S&P500"}

	ProcessPayment(cc)
	ProcessPayment(bc)
	ProcessPayment(ba)
	ProcessPayment(st)
}
