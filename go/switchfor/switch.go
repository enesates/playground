package main

import "fmt"

func calculator(a, b float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0.0, fmt.Errorf("unknown operation: %s", operation)
	}
}

func main() {
	if val, err := calculator(30.0, 50.0, "+"); err != nil {
		fmt.Printf("%.2f + %.2f = %.2f\n", 30.0, 50.0, val)
	} else {
		fmt.Println(err)
	}
}
