package main

import (
	"fmt"
)

func main() {
	var numbers string
	fmt.Print("Ввод номера (четного):\n")
	fmt.Scanf("%s", &numbers)
	fmt.Println(isLucky(numbers))
}

func isLucky(k string) bool {
	var (
		sum1 int = 0
		sum2 int = 0
		len  int = 0
	)
	for range k {
		len++
	}
	r := len / 2
	i1 := k[:r]
	i2 := k[r:]
	for i := range i1 {
		sum1 += i
	}
	for j := range i2 {
		sum2 += j
	}
	if sum1 == sum2 {
		return true
	} else {
		return false
	}

}
