package main

import "fmt"

func main() {

	var word string
	fmt.Println("Введите слово")
	fmt.Scanf("%s\n", &word)

	sample := word
	r := []rune(sample)
	var word1 []rune
	for i := len(r) - 1; i >= 0; i-- {
		word1 = append(word1, r[i])
	}

	s := string(word1)
	//fmt.Println(s)

	result := s == word

	//fmt.Println("\nResult: ", result)

	if result == true {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Не палиндром")
	}

}
