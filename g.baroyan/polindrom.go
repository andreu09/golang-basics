package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var stroka string
	var check bool
	fmt.Println("Введите строку для проверки на палиндром: ")
	fmt.Scan(&stroka)

	size := utf8.RuneCountInString(stroka)

	massiv := make([]string, size)

	j := size - 1
	for i := 0; i < size; i++ {
		massiv[i] = string(stroka[j])
		j--
	}

	for i := 0; i < size; i++ {
		if massiv[i] != string(stroka[i]) {
			check = false
			break
		} else {
			check = true
		}
	}
	if check == true {
		fmt.Println("Cтрока является полиндромом!")
	} else {
		fmt.Println("Cтрока не является полиндромом!")
	}

}
