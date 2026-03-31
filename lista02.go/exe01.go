package main

import f "fmt"

func main() {

	var n int

	f.Println("Digite um número inteiro: ")
	f.Scan(&n)

	if n%2 == 0 {
		f.Println("O número é par!!!")
	} else {
		f.Println("O número é ímpar!!!")
	}
}
