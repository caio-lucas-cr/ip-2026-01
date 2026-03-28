package main

import f "fmt"

func main() {

	//Algoritmo "Somatório Simples"

	var n, i int
	var S float32

	//Entrada
	f.Println("Digite um número inteiro: ")
	f.Scan(&n)

	//Saída
	if n > 1 {
		for i = 1; i <= n; i++ {
			S = S + (1 / float32(i))
		}
	} else {
		f.Println("Número inválido")
	}

	f.Printf("A soma é: %.6f", S)

}
