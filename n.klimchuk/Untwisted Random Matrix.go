package main

import (
	"fmt"
	"math/rand"
	"time"
)

func some_matrix(input int, some_matrix [][]int) {
	i := input / 2
	j := input / 2

	n := 0

	for true {
		n++

		for k := 0; k < n; k++ {
			fmt.Print(some_matrix[i][j])
			j--
		}

		if n == input {
			break
		}

		for k := 0; k < n; k++ {
			fmt.Print(some_matrix[i][j])
			i++
		}

		n++

		for k := 0; k < n; k++ {
			fmt.Print(some_matrix[i][j])
			j++
		}

		for k := 0; k < n; k++ {
			fmt.Print(some_matrix[i][j])
			i--
		}
	}
}

func main() {
	var n int

	fmt.Printf("Enter the initial dimension of the matrix... ")
	fmt.Scan(&n)
	fmt.Println()

	dimension := 2*n - 1

	matrix := make([][]int, dimension)

	for i := range matrix {
		matrix[i] = make([]int, dimension)
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			matrix[i][j] = rand.Intn(10)
		}
	}

	fmt.Println(matrix)
	fmt.Println()

	some_matrix(dimension, matrix)
}
