package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main(){
	// считываем ввод пользователя
	fmt.Print("Введите строку для проверки, является ли она палиндромом: ")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	
	// Собираем из строки все, что не является пробелом или переносом строки, переводим в нижний регистр
	var text string
	for _, i := range line{
		if string(i) != " " && string(i) != "\n" {
			text += strings.ToLower(string(i))
		}
	}
	
	// Конвертируем строку в список рун и переворачиваем(записываем в обратном порядке элементы списка в новый список)
	r := []rune(text)
	var res []rune
	for i := len(r) - 1; i>=0; i--{
		res = append(res, r[i])
	}

	// выводим результат
	if text == string(res){
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}
}