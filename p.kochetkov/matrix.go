package main

import (
	"fmt"
	"math/rand"
)


func randon_matrix(n int) [][]int {
	A := make([][]int, n)

	for i := 0; i < len(A); i++ {
		StrokLen := len(A)
		A[i] = make([]int, StrokLen)
		for j := 0; j < StrokLen; j++ {
			A[i][j] = rand.Intn(10)
		}
	}
	return A

}


func revers(s [][]int) [][]int {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}


func transpose(M [][]int) [][]int {

	m := len(M)    
	n := len(M[0]) 
	
	T := make([][]int, n)
	for i := 0; i < n; i++ {
		T[i] = make([]int, m)
	}

	
	var row []int
	for i := 0; i < n; i++ {
		row = T[i]
		for j := 0; j < m; j++ {
			row[j] = M[j][i]
		}
	}

	return T
}


func rotation(A [][]int) [][]int {
	var sd [][]int

	sd = A[1:]
	sd = transpose(sd)
	sd = revers(sd)

	return sd
}

func main() {

	const n int = 5 

	var dom [][]int 

	S := randon_matrix(n) 

	dom = append(dom, S[0]) 

	
	for i := 0; i < len(S); i++ {
		fmt.Println(S[i])
	}

	
	for len(S) >= 2 {
		S = rotation(S)
		dom = append(dom, S[0])
	}

	for i := len(dom) - 1; i >= 0; i-- {
		for j := len(dom[i]) - 1; j >= 0; j-- {
			fmt.Print(dom[i][j])
		}
	}
}

