package main

import (
 "fmt"
 "unicode/utf8"
)

func main() {

 var inputStr string
 fmt.Println("Введите строку для проверки на палиндром: ")
 fmt.Scan(&inputStr)

 size := utf8.RuneCountInString(inputStr)

 arr := make([]string, size)

 j := size - 1
 for i := 0; i < size; i++ {
  arr[i] = string(inputStr[j])
  j--
 }

 for i := 0; i < size; i++ {
  if arr[i] != string(inputStr[i]) {
   fmt.Println("Строка не палиндром")
   break
  } else if i == size-1 {
   fmt.Println("Строка палиндром")

  }
 }

}