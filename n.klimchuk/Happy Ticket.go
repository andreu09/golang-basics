package main

import (
	"fmt"
	"strconv"
)

func some_ticket(input string) bool {
	check_ticket := len(input)

	for check_ticket%2 != 0 {
		fmt.Println("Hey, that's the wrong number!")

		fmt.Printf("Enter the correct number... ")
		fmt.Scan(&input)
		fmt.Println()

		check_ticket = len(input)
	}

	half_line := check_ticket / 2

	part_one := 0
	part_two := 0

	some_number, some_error := strconv.Atoi(input)

	if some_error != nil {
		fmt.Println("Something went wrong!")
	}

	some_array := make([]int, check_ticket)

	for i := 0; i < check_ticket; i++ {
		some_array[i] = some_number % 10
		some_number /= 10
	}

	for i := 0; i < check_ticket; i++ {
		if i <= half_line-1 {
			part_one += some_array[i]
		} else {
			part_two += some_array[i]
		}
	}

	if part_one != part_two {
		return false
	} else {
		return true
	}
}

func main() {
	var some_string string

	fmt.Printf("Enter the ticket number for verification... ")
	fmt.Scan(&some_string)
	fmt.Println()

	fmt.Println(some_ticket(some_string))
}
