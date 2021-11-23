package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Дополнительное задание\nВыполнил Шмаков А.Ю.")
	var n int
	k := 0
	z := 0
	fmt.Println("Введите кол-во элементов массива: ")
	fmt.Scanf("%d\n", &n)
	mas := make([]int, n)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		mas[i] = rand.Intn(100)
	}

	fmt.Println("Исходный массив: ")

	fmt.Print(mas)

	for i := 0; i < (n - 1); i++ {
		// Если текущий элемент больше предыдущего
		if mas[i] > mas[i+1] {
			// Накапливаем шаг, он равен разнице следующего элемена и предыдущего минус 1 по модулю
			k += int(math.Abs(float64(mas[i+1] - mas[i] - 1)))
			z = int(math.Abs(float64(mas[i+1] - mas[i] - 1)))
			// Прибавляем к изменяемому элементу посчитанную разницу
			mas[i+1] = mas[i+1] + z

		}

		// Частный случай если предыдущий элемент раверн следующему
		// Все так же как и выше, только + 1
		if mas[i] == mas[i+1] {
			k += 1
			z = 1
			mas[i+1] = mas[i+1] + z

		}

	}

	fmt.Print("\nВаша новая возрастающая последовательность:\n")
	fmt.Println(mas)
	fmt.Println("Минимальное кол-во шагов:", k)

}
