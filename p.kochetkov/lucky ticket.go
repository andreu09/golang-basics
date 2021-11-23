package main

import (
 "fmt"
 "strconv"
 "unicode/utf8"
)

func main() {

 var inputStr string
 fmt.Println("Введите номер билета для его проверки на счастливость: ")
 fmt.Scan(&inputStr)

 size := utf8.RuneCountInString(inputStr)

 for size%2 != 0 {
  fmt.Println("Введите четное кол-во символов: ")
  fmt.Scan(&inputStr)
  size = utf8.RuneCountInString(inputStr)

 }

 sum1 := 0
 sum2 := 0

 num, err := strconv.Atoi(inputStr)
 if err != nil {
  panic(err)
 }

 arr := make([]int, size)

 for i := 0; i < size; i++ {
  arr[i] = num % 10
  num /= 10
 }

 for i := 0; i < size; i++ {
  if i <= size/2-1 {
   sum1 += arr[i]
  } else {
   sum2 += arr[i]

  }
 }

 if sum1 == sum2 {
  fmt.Println("Ваш билет счастливый")
 } else {
  fmt.Println("Ваш билет не счастливый")

 }

}