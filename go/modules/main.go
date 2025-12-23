package main

import (
	"fmt"

	"github.com/enesates/playground/go/externallibrary/mathlib/advanced"
	"github.com/enesates/playground/go/externallibrary/mathlib/operations"
	"github.com/enesates/playground/go/externallibrary/queue"
)

func main() {
	add := operations.Addition(2, 3)
	fmt.Printf("2 + 3 = %.2f\n", add)

	///////////////////

	pow, err := advanced.Power(3, 4)
	if err != nil {
		fmt.Printf("3 ** 4 = %+v\n", err)
	} else {
		fmt.Printf("3 ** 4 = %d\n", pow)
	}

	pow2, err := advanced.Power(3, -1)
	if err != nil {
		fmt.Printf("3 ** -1 = %+v\n", err)
	} else {
		fmt.Printf("3 ** -1 = %d\n", pow2)
	}

	///////////////////

	q := queue.Push(5)
	fmt.Println(q)

	q = queue.Push(6)
	fmt.Println(q)

	q = queue.Pop()
	fmt.Println(q)
}
