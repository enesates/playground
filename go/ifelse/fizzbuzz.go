package main

import "fmt"

func fizzbuzz(n int) {
	if n%15 == 0 {
		fmt.Print("FizzBuzz")
	} else if n%3 == 0 {
		fmt.Print("Fizz")
	} else if n%5 == 0 {
		fmt.Print("Buzz")
	} else {
		fmt.Println("Nothing")
	}
}

func main() {
	fizzbuzz(15)
}
