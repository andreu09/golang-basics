package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {

	var inpStr string
tryAgain:
	fmt.Println("Введите номер билета, чтобы узнать является ли он счастливым: ")
	fmt.Scan(&inpStr)

	//определяем количество символов во введенном пользователе номере билета
	size := utf8.RuneCountInString(inpStr)

	//если символов нечетное количество требуем ввести снова, пока оно не будет четным
	for size%2 != 0 {
		fmt.Println("Введите четное кол-во символов: ")
		fmt.Scan(&inpStr)
		size = utf8.RuneCountInString(inpStr)
	}

	//конвертируем строку в число, проворяем успешна ли конвертация
	num, err := strconv.Atoi(inpStr)
	if err != nil {
		fmt.Println("Ошибка! Вы ввели не число!")
		goto tryAgain
	}

	//создаем массив на количество элементов равное количеству символов в билете
	number := make([]int, size)

	//заполняем массив числами
	for i := 0; i < size; i++ {
		number[i] = num % 10
		num /= 10
	}

	//создаем две переменные для сравнения суммы правой и левой части билета
	firstSum := 0
	secSum := 0

	//считаем сумму правой и левой части билета
	for i := 0; i < size; i++ {
		if i <= size/2-1 {
			firstSum += number[i]
		} else {
			secSum += number[i]
		}
	}

	//выводим результат
	if firstSum == secSum {
		fmt.Println("Урааа! Ваш билет счастливый")
	} else {
		fmt.Println("Увы... Ваш билет не счастливый")
	}

}
