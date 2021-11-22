package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var inpStr string
	fmt.Println("Введите строку для проверки на палиндром: ")
	fmt.Scan(&inpStr)

	//определяем длину введенной строки
	size := utf8.RuneCountInString(inpStr)

	//создаем массив нв количество элементов равной длине введенной строки
	word := make([]string, size)

	//заполняем созданный массив строкой в обратном её порядке
	j := size - 1
	for i := 0; i < size; i++ {
		word[i] = string(inpStr[j])
		j--
	}

	//сравниваем каждый элемент массива с элементом строки
	for i := 0; i < size; i++ {
		if word[i] != string(inpStr[i]) { //если хоть один символ отличатся выходим из цикла и выводим ответ
			fmt.Println("Строка ", inpStr, " не палиндром")
			break
		} else if i == size-1 { //если проитерировались по всей длине и все символы равны - строка палиндром
			fmt.Println("Строка ", inpStr, " палиндром")
		}
	}

}
