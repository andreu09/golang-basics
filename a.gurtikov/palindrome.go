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

	if size == 1 {
		fmt.Println("Строка ", inpStr, " палиндром!")
	} else if size == 2 {
		if inpStr[0] == inpStr[1] {
			fmt.Println("Строка ", inpStr, " палиндром!")
		} else {
			fmt.Println("Строка ", inpStr, " не палиндром!")
		}
	} else {
		i := 0
		j := size - 1
		delta := 0
		for {
			if inpStr[i+delta] == inpStr[j-delta] {
				j--
				i++
				delta++

				if i == size/2 {
					fmt.Println("Строка ", inpStr, " палиндром!")
					break
				}

			} else {
				fmt.Println("Строка ", inpStr, " не палиндром!")
				break
			}

		}
	}

}
