// перенос решения задачи с языка С++

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// функция создания и заполнения матрицы псевдослучайными числами
func createMatrix(n int) ([][]int, int) {

	size := 2*n - 1

	matrix := make([][]int, size)

	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = rand.Intn(9)
		}
	}

	return matrix, size

}

//функция вывода созданное матрицы
func printMatrix(matrix [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

// функция пользовательского ввода
func userInput() int {
	var n int
	fmt.Println("Введите n для создания квадратной матрицы 2n-1 x 2n-1")
	fmt.Scan(&n)
	for n <= 1 {
		fmt.Println("Ошибка! Введите число больше чем 1")
		fmt.Scan(&n)
	}
	return n
}

// функция вывода матрицы по спирали
func printSpiral(matrix [][]int, n int) {

	step := -1
	iter := 1
	fmt.Print("Вывод по спирали: ", matrix[n/2][n/2], " ")
	i := n / 2
	j := n / 2 //нашли центр матрицы
	k := 0
	l := 0   // итераторы
	out := 1 //счетчик

link:
	for out != n*n { //пока количество выведенных символов не равно кол-ву всех элементов в матрице

		for ; k < iter; k++ {
			if j+step < n && j+step >= 0 { //если совершая шаг мы не покидаем пределы матрицы
				fmt.Print(matrix[i][j+step], " ") //выводим число в консоль
				out++                             // увеличиваем счетчик отвечающий за выход
				if out == n*n {
					goto link //если счетчик равен количеству чисел в матрице идем проверять условие главного цикла и покидаем его
				}
			}

			if iter > 1 && k+1 < iter {
				if step > 0 {
					step++
				} else {
					step--
				}
			}

		}

		j = j + step
		k = 0

		if step > 0 {
			step = 1
		} else {
			step = -1
		}

		step *= -1

		for ; l < iter; l++ {

			if i+step < n && i+step >= 0 {

				fmt.Print(matrix[i+step][j], " ")
				out++
				if out == n*n {
					goto link
				}
			}
			if iter > 1 && l+1 < iter {
				if step > 0 {
					step++
				} else {
					step--
				}
			}

		}
		i = i + step

		if step > 0 {
			step = 1
		} else {
			step = -1
		}

		iter++
		l = 0

	}

}

func main() {

	n := userInput()
	matrix, size := createMatrix(n)
	printMatrix(matrix, size)
	printSpiral(matrix, size)

}
