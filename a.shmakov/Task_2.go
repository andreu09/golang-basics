package main

import (
	"fmt"
)

func main() {
	var str string
	flag := true
	var lenStr int

	fmt.Println("Задание 2\nВыполнил Шмаков А.Ю.")
	fmt.Println("Введите строку: ")
	fmt.Scanf("%s\n", &str)
	lenStr = len(str)
	for i := 0; i < (lenStr / 2); i++ {
		if str[i] != str[lenStr-i-1] {
			flag = false
		}
	}

	fmt.Println("Результат: ", flag)
}
