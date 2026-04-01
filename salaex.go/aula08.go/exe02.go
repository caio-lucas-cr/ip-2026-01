package main

import f "fmt"

func main() {

	var n [3]int
	var Mn int

	for i := 0; i < len(n); i++ {
		f.Printf("Digite o %d° número: \n", i+1)
		f.Scan(&n[i])
	}

	Mn = maiorNumero(n)

	f.Printf("O maior número é o %d \n", Mn)
}

func maiorNumero(n [3]int) int {

	var m int

	if n[0] > n[1] && n[0] > n[2] {
		m = n[0]
	} else if n[1] > n[0] && n[1] > n[2] {
		m = n[1]
	} else if n[2] > n[1] && n[2] > n[0] {
		m = n[2]
	}
	return m
}
