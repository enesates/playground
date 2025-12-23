package main

import "fmt"

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func isPalindrome2(s string) bool {
	r := []rune(s)

	var rec func(i, j int) bool
	rec = func(i, j int) bool {
		if i >= j {
			return true
		}
		if r[i] != r[j] {
			return false
		}
		return rec(i+1, j-1)
	}

	return rec(0, len(r)-1)
}

func main() {
	fmt.Println("is enes Palindrome?:", isPalindrome("enes"))
	fmt.Println("is rentner Palindrome?:", isPalindrome("rentner"))
}
