package main

import "fmt"

func main() {
	productsAndPrices := map[string]int{
		"Book": 10,
		"Ball": 20,
		"Bag":  30,
	}

	productsAndPrices["Book"] = 12

	toyPrice, ok := productsAndPrices["Toy"]

	if ok {
		fmt.Println("Toy price:", toyPrice)
	} else {
		fmt.Println("No Toy!")
	}

	delete(productsAndPrices, "Bag")

	fmt.Println(productsAndPrices)
}
