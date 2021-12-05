package main

import "fmt"

func some_palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var some_string string

	fmt.Printf("Enter a string to check for palindrome... ")
	fmt.Scan(&some_string)
	fmt.Println()

	fmt.Println(some_palindrome(some_string))
}
