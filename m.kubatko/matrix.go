package main
import (
	"fmt"
)
func main(){
	// считываем ввод пользователя
	var in int
	fmt.Print("Введите число: ")
	fmt.Scan(&in) 
	
	s := 2*in-1
	c := 0
	var row []int
	// заполняем матрицу числами
	matrix := make([][]int , s)
	for num :=1; num <= (s*s); num++ {
		row = append(row, num)

		if len(row) == s {			
		    matrix[c] = row
			row = nil
			c ++
			continue
		}
	}
	// правила движения по спирали
	move := [] map [string] int {
		{"i": 0, "j": -1},
		{"i": 1, "j": 0},
		{"i": 0, "j": 1},
		{"i": -1, "j": 0},
	}

	count := 0
    len_step  := 1
	
	// если ввели 1 выводим ее
	if len(matrix)==1{
		fmt.Println(matrix[0][0])
	}else{
		// находим середину матрицы и начинаем заполнять срез
		mid := s/2
		var sp []int
		i := mid
		j := mid
		sp = append(sp, matrix[i][j])
		
		f := 1
		// устанавливаем флаг для прерывания цикла
		for f == 1{
			move_indx := count % 4
			mi := move[move_indx]["i"]
			mj := move[move_indx]["j"]
			// вычисляем нужные позиции чисел, заполняем срез
			for n := 0; n < len_step; n++{
				i += mi
				j += mj
				sp = append(sp, matrix[i][j])
				// начало матрицы и конец цикла
				if i == 0 && j == 0{
					f = 0
					break
				}
			}
			// после нечетного шага (т.е.каждые два) увеличиваем длину отрезка
			len_step = len_step+(count%2)
			count++
		}
		// вывод в строку содержимого среза
		for _, val := range sp{
			fmt.Print(val, " ")
		}
		fmt.Println()
	}	
}