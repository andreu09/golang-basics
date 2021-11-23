package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	Up = iota
	Right
	Down
	Left
)

func main() {
	n := get_data_from_input()
	condition_of_matrix(n)
}

func get_data_from_input() uint16 {
	var input string
	fmt.Println("Введите n")
	for {
		fmt.Scanln(&input)
		uint_input, err := strconv.ParseUint(input, 10, 16)
		if err != nil || uint_input == 0 {
			fmt.Println("Введите целочисленное значение")
			continue
		}

		return uint16(uint_input)
	}
}

func condition_of_matrix(n uint16) {
	var size uint16 = 2*n - 1
	rand_matrix := make([][]int, size)
	for i := 0; i < int(size); i++ {
		rand_matrix[i] = make([]int, size)
	}
	generate_matrix(rand_matrix)
	print_matrix(rand_matrix)
	spiral(rand_matrix)

}

func generate_matrix(rand_matrix [][]int) {
	for i, innerArray := range rand_matrix {
		for j := range innerArray {
			rand_matrix[i][j] = rand.Intn(9)
		}
	}
}

func print_matrix(rand_matrix [][]int) {
	for i := range rand_matrix {
		fmt.Println(rand_matrix[i])
	}
}

func spiral(matrix [][]int) {
	output := make([]int, len(matrix)*len(matrix[0]))
	direction := Right
	col, minRow, maxRow := 0, 1, len(matrix)-1
	row, minCol, maxCol := 0, 0, len(matrix[0])-1
	for i := 0; i < len(output); i++ {
		output[i] = matrix[row][col]
		switch direction {
		case Right:
			if col == maxCol {
				direction = Down
				maxCol--
			}
		case Down:
			if row == maxRow {
				direction = Left
				maxRow--
			}
		case Left:
			if col == minCol {
				direction = Up
				minCol++
			}
		case Up:
			if row == minRow {
				direction = Right
				minRow++
			}
		}
		switch direction {
		case Right:
			col++
		case Down:
			row++
		case Left:
			col--
		case Up:
			row--
		}
	}
	for i, j := 0, len(output)-1; i < j; i, j = i+1, j-1 {
		output[i], output[j] = output[j], output[i]
	}
	fmt.Println(output)
}
