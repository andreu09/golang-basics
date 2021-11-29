package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generate(randMatrix [][]int) {
	for i, innerArray := range randMatrix {
		for j := range innerArray {
			randMatrix[i][j] = rand.Intn(100)

		}
	}
}

func main() {
	var n, k, a1 int
	fmt.Println("Введите число")
	fmt.Scan(&n)
	k = 2*n - 1

	randMatrix := make([][]int, k)

	for i := 0; i < k; i++ {
		randMatrix[i] = make([]int, k)
	}

	generate(randMatrix)
	//fmt.Println(randMatrix)

	for j := 0; j < k; j++ {
		fmt.Println(randMatrix[j][0:k])

	}

	a2 := make([]int, 0)
	a2 = append(a2, randMatrix[n-1][n-1])
	for j := 1; j < n; j++ {

		for i := 0; i < j+1; i++ {
			a1 = randMatrix[n-1+i][n-j-1]
			a2 = append(a2, a1)
		}
		for i := 0; i < 2*j; i++ {
			a1 = randMatrix[n-1+j][n-j+i]
			a2 = append(a2, a1)
		}
		for i := 0; i < 2*j; i++ {
			a1 = randMatrix[n+j-2-i][n+j-1]
			a2 = append(a2, a1)
		}
		for i := 0; i < 2*j; i++ {
			a1 = randMatrix[n-j-1][n-i+j-2]
			a2 = append(a2, a1)
		}
	}
	fmt.Println(a2)

}
