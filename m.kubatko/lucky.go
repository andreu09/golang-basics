package main
import (
	"fmt"
	"strconv"
)

func main(){
	// считываем ввод пользователя
	var in string
	fmt.Print("Введите номер билета, чтобы проверить счастливый ли он: ")
	fmt.Scan(&in) 
	// проверяем, что введены цифры
	if _, err := strconv.Atoi(in); err == nil {
		// проверяем, что введено четное количество цифр
		if len(in)%2 == 0{
			// переводим введеную строку в срез интов
			var numln []int
			for i := 0; i<=len(in)-1; i++ {
				n,_ := strconv.Atoi(string(in[i]))
				numln = append(numln, n)
			}
			// определяем середину среза
			mid := len(numln)/2

			var num1 int
			var num2 int
			// ссумируем содержимое одной половины и второй половины среза.
			for i := 0; i <= len(numln)-1; i++{
				if i < mid{
					num1 += numln[i]
				}else{
					num2 += numln[i]					
				}				
			}
			// сравниваем содержимое половин
			if num1 == num2 {
				fmt.Println(true)
			}else{
				fmt.Println(false)
			}
		}else{
			fmt.Println("Ошибка! Вы ввели нечетное число.")
		}		
	}else{
		fmt.Println("Ошибка! Вы указали буквы вместо цифр")
	}

}