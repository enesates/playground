package main

import (
	"fmt"
)

func division(num1, num2 float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ERROR:", r)
		}
	}()

	if num2 == 0 {
		panic("DIVISION BY ZERO")
	}

	fmt.Printf("%.2f / %.2f = %2f", num1, num2, num1/num2)
}

func main() {
	num1, num2 := 5.0, 0.0
	division(num1, num2)
}
