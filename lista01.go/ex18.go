package main

import f "fmt"

func main() {

	//Algoritmo "Soma de progressão aritmética"

	var a1, r, N, i, S int

	//Entrada
	f.Println("Digite o número incial da PA, a razão e a quantidade de termos, respectivamente: ")
	f.Scan(&a1, &r, &N)

	//Processamento
	for i = 1; i <= N; i++ {
		S = S + (a1 + (i-1)*r)
	}

	//Saída
	f.Println("A SOMA É: ", S)

}
