package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {

	var Str string
	fmt.Println("Введите номер билета")
	fmt.Scan(&Str)

	size := utf8.RuneCountInString(Str)

	if size%2 != 0 {
		fmt.Println("Введите четное кол-во символов: ")
		fmt.Scan(&Str)
		size = utf8.RuneCountInString(Str)

	}

	sum1 := 0
	sum2 := 0

	num, e := strconv.Atoi(Str)
	if e != nil {
		fmt.Println(e)
	}

	massiv := make([]int, size)

	for i := 0; i < size; i++ {
		massiv[i] = num % 10
		num /= 10
	}

	for i := 0; i < size/2; i++ {
		sum1 += massiv[i]
	}
	for i := size / 2; i < size; i++ {
		sum2 += massiv[i]
	}

	if sum1 == sum2 {
		fmt.Println("Билет является счастливым")
	} else {
		fmt.Println("Билет не является счастливым")

	}

}
