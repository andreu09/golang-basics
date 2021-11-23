package main

import "fmt"

func main() {
	var n int          // Размерность матрицы
	a := 1             // Значение нулевого элемент матрицы
	var stepLB int = 1 // Начальное кол-во шагов по матрице влево и вниз
	var stepRT int = 2 // Начальное кол-во шагов по матрице вправо и вверх
	var step int = 0   // Счетчик шагов на каждой итерации
	fmt.Println("Задание 1\nВыполнил Шмаков А.Ю.")
	fmt.Println("Введите значение n = ")
	fmt.Scanf("%d\n", &n)
	n = (2 * n) - 1
	var initI int = (n / 2) // Значение строки элемента, с которого начинается итерация
	var initJ int = (n / 2) // Значение столбца элемента, с которого начинается итерация
	// Матрица элементов
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
	}

	fmt.Println("Размерность матрицы: ", n, "x", n)

	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {
			arr[i][j] = a
			a++
		}
	}

	fmt.Println("Введенная матрица:")

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Print("\n")
	}

	fmt.Println("Результат:")

	for {

		// Цикл прохода по матрице влево
		for j := initJ; step < stepLB; {
			fmt.Print(arr[initI][j], " ")
			step++
			j--
			initJ = j
		}

		// Алгоритим заканчивается, если кол-во шагов влево стало равно размерности матрицы
		if stepLB == n {
			break
		}

		step = 0

		// Цикл прохода по матрице вниз
		for i := initI; step < stepLB; {
			fmt.Print(arr[i][initJ], " ")
			step++
			i++
			initI = i
		}

		step = 0

		// Цикл прохода по матрице вправо
		for j := initJ; step < stepRT; {
			fmt.Print(arr[initI][j], " ")
			step++
			j++
			initJ = j
		}

		step = 0

		// Цикл прохода по матрице вверх
		for i := initI; step < stepRT; {
			fmt.Print(arr[i][initJ], " ")
			step++
			i--
			initI = i
		}

		step = 0
		// Увеличиваем кол-во шагов в стороны по матрице
		stepLB += 2
		stepRT += 2
	}

	fmt.Print("\n")

}
