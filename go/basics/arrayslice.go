package main

import "fmt"

func main() {
	var shoppingList []string
	shoppingList = append(shoppingList, "Milk", "Bread")

	extras := []string{"Egg", "Cheese"}
	shoppingList = append(shoppingList, extras...)

	fmt.Println("Length:", len(shoppingList))
	fmt.Println("Capacity:", cap(shoppingList))

	shoppingList = append(shoppingList[:1], shoppingList[2:]...)

	fmt.Println("New Shopping List:", shoppingList)
}
