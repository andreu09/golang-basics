package main

import (
	"fmt"
)

func main() {
	var stroka string
	fmt.Print("Ввод строки:\n")
	fmt.Scanf("%s", &stroka)
	fmt.Print(isPolindrom(stroka))
}

func isPolindrom(k string) bool {
	var (
		len int
	)
	for range k {
		len++
	}
	if len == 1 {
		return true
	}
	for i := 0; i < len/2; i++ {
		if k[i] != k[len-i-1] {
			return false
		}
	}
	return true
}
