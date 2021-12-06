package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		step       int = 0
		max_value  int
		next_value int
		len        int
	)
	rand.Seed(time.Now().UnixNano())
	massiv := make([]int, 0, 11)
	for i := 0; i < rand.Intn(12); i++ {
		massiv = append(massiv, rand.Intn(100))
	}
	fmt.Print(massiv)
	for range massiv {
		len++
	}
	for k := 0; k < len+1; k++ {
		fmt.Print(k)
		if k+1 == len {
			break
		}
		max_value = massiv[k]
		next_value = massiv[k+1]
		for next_value <= max_value {
			next_value += 1
			step += 1
		}
		massiv[k+1] = next_value
	}
	fmt.Print(massiv)
	fmt.Print(step)
}
