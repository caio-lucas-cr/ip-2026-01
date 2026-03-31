package main

import f "fmt"

func main() {

	var A, B int

	f.Println("Digite um número A:")
	f.Scan(&A)

	f.Println("Digite um número B:")
	f.Scan(&B)

	if A%B == 0 {
		f.Println("O número", A, "É divisível por", B)
	} else {
		f.Println("O número", A, "NÃO É divisível por", B)
	}
}