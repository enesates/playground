package main

import "fmt"

func checkType(val interface{}) {
	_, ok := val.(string)
	if ok {
		fmt.Println(val, "is a string")
	}

	_, ok = val.(int)
	if ok {
		fmt.Println(val, "is an integer")
	}

	fmt.Println(val, "neither a number nor a letter")
}

func main() {
	checkType("enes")
	checkType(2)
}
