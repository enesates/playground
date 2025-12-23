package main

import "fmt"

func main() {
	var celsius float64
	const factor = 1.8
	const offset = 32.0

	fmt.Print("Enter Celsius Value: ")
	_, _ = fmt.Scanln(&celsius)

	fahrenheit := celsius*factor + offset

	fmt.Printf("Celsius: %.2fC\n", celsius)
	fmt.Printf("Fahrenheit: %.2fF\n", fahrenheit)
}
