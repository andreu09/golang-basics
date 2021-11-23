package main

import (
	"fmt"
)

func main() {
	var str string

	fmt.Println("Задание 3\nВыполнил Шмаков А.Ю.")
	fmt.Println("Ведите номер билета: ")
	fmt.Scanf("%s\n", &str)
	lenStr := (len(str) / 2)
	var sum1 int = 0
	var sum2 int = 0
	if len(str)%2 == 0 {

		str1 := str[lenStr:]
		str2 := str[:lenStr]
		for runeValue := range str1 {
			sum1 += int(str1[runeValue])
		}

		for runeValue := range str2 {
			sum2 += int(str2[runeValue])
		}

		if sum1 == sum2 {
			fmt.Println("Счастливый билет!")
		} else {
			fmt.Println("Не повезло")
		}

	} else {

		fmt.Println("Номер билета должен содержать четное кол-во знаков!")

	}

}
