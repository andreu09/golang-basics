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
	num_of_thicket := read_input()
	check_ticket(num_of_thicket)
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
		regexp := regexp.MustCompile("[0-9]")

		if !regexp.MatchString(input) {
			fmt.Println("Повторите попытку. Значение должно быть целым числом")
			continue
		} else if len(input) < 2 {
			fmt.Println("Слишком короткая число. Повторите попытку")
			continue
		} else if len(input)%2 != 0 {
			fmt.Println("Введите четное количество символов")
			continue
		}

		return input
	}
}

func check_ticket(num_of_thicket string) {
	center_of_string := len(num_of_thicket) / 2

	first_part := num_of_thicket[center_of_string:]
	second_part := num_of_thicket[:center_of_string]

	first_sum := 0
	second_sum := 0

	for first_rune := range first_part {
		first_sum += int(first_part[first_rune])
	}

	for second_rune := range second_part {
		second_sum += int(second_part[second_rune])
	}

	if first_sum == second_sum {
		fmt.Println("Удача!")
	} else {
		fmt.Println("В следующий раз повезет")
	}

}
