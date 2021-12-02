package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Введите значение")
	input := read_input()
	check_palindrome(input)
}

func read_input() string {
	for {

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Возникла проблема при попытке прочитать буфер, повторите попытку", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")
		input = strings.ReplaceAll(input, " ", "")
		regexp := regexp.MustCompile("[a-zA-Z0-9]")

		if !regexp.MatchString(input) {
			fmt.Println("Повторите попытку. Значение должно быть числом или символом английского алфовита")
			continue
		} else if len(input) < 2 {
			fmt.Println("Слишком короткая строка. Повторите попытку")
			continue
		}

		return input
	}
}

func check_palindrome(input string) {
	lower_input := strings.ToLower(input)

	runes := []rune(lower_input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	rev_string := string(runes)

	if lower_input == rev_string {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}
