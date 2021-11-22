package main

import (
	"fmt"
	"math/rand"
	"strconv"
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

//функция вывода созданной матрицы
func printMatrix(matrix [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

// функция пользовательского ввода числа n
func userInput() int {
	var inpStr string
tryAgain:
	fmt.Println("Введите n для создания квадратной матрицы размером 2n-1 x 2n-1")
	fmt.Scan(&inpStr)

	//конвертируем строку в число, проворяем успешна ли конвертация
	n, err := strconv.Atoi(inpStr)

	if err != nil {
		fmt.Println("Внимание! Вы ввели не число!")
		goto tryAgain
	}

	for n <= 1 {
		fmt.Println("Ошибка! Введите число больше чем 1")
		fmt.Scan(&n)
	}
	return n
}

// функция вывода матрицы по спирали
func printSpiral(matrix [][]int, n int) {

	center := n / 2        //нашли центр матрицы (начало пути)
	x, y := center, center //координаты для перемещения по матрице

	var left, down, right, up bool // направление шагов
	left = true

	step := 1      //величина шага
	stepCount := 0 //счетчик шагов
	nextStep := 1  //следующая величина шага
	work := true   //условие работы главного цикла
	count := 1     //количество выведенных чисел

	fmt.Print("Вывод по спирали: ", matrix[x][y], " ")

link:
	for work {

		switch {
		case left: //  шагаем влево

			for i := 0; i < step; i++ {
				y--
				fmt.Print(matrix[x][y], " ")
				stepCount++
				count++
				if x == 0 && y == 0 {
					work = false
					goto link
				}
			}

			if stepCount == nextStep*2 {
				step++
				nextStep++
				stepCount = 0

			}
			left = false
			down = true

		case down: //  шагаем вниз

			for i := 0; i < step; i++ {
				x++
				fmt.Print(matrix[x][y], " ")
				count++
				stepCount++
			}

			if stepCount == nextStep*2 {
				step++
				nextStep++
				stepCount = 0
			}

			down = false
			right = true

		case right: //  шагаем вправо

			for i := 0; i < step; i++ {
				y++
				fmt.Print(matrix[x][y], " ")
				count++
				stepCount++
			}

			if stepCount == nextStep*2 {
				step++
				nextStep++
				stepCount = 0
			}

			right = false
			up = true

		case up: //  шагаем вверх

			for i := 0; i < step; i++ {
				x--
				fmt.Print(matrix[x][y], " ")
				count++
				stepCount++
			}

			if stepCount == nextStep*2 {
				step++
				nextStep++
				stepCount = 0
			}
			up = false
			left = true
		}
	}

	fmt.Println()
	fmt.Println("Выведено чисел: ", count)

}

func main() {

	n := userInput()
	matrix, size := createMatrix(n)
	printMatrix(matrix, size)
	printSpiral(matrix, size)

}
