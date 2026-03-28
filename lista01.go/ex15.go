package main

import f "fmt"

func main() {
	//Algoritmo "Quadrado de pares"

	var N, k int

	//Entrada
	f.Println("Digite um número: ")
	f.Scan(&N)

	//Saída
	if N <= 5 || N >= 200 {
		f.Println("Não é possível realizar a operação")
	}

	for i := 1; i <= N; i++ {
		if i%2 == 0 && N > 5 && N < 2000 {
			k = i * i
			f.Println(i, "² =", k)
		}
	}

}
