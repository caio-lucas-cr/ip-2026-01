package main

import f "fmt"

func main() {

	var n int

	f.Println("Digite um número inteiro: ")
	f.Scan(&n)

	if n%5 == 0 {
		f.Println("O número É divisível por 5")
	} else {
		f.Println("O número NÃO É divisível por 5")
	}
}