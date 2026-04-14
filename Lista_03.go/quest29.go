package main

import f "fmt"

func main() {

	var N, S int

	f.Println("O somatório de todos os números inteiros 1 até N é dado por: ")
	f.Println("Insira o valor de N: ")
	f.Scan(&N)

	if N <= 1 {
		f.Println("O número inserido é inválido")
		return
	}

	S = (1 + N) * N / 2

	f.Printf("\nS = %d", S)

}
