package main

import f "fmt"

func main() {

	// Algoritmo "Composição Inteira"

	var n1, n2, n3, N, Q int

	//Entrada
	f.Println("Digite o número 1: ")
	f.Scan(&n1)

	f.Println("Digite o número 2: ")
	f.Scan(&n2)

	f.Println("Digite o número 3: ")
	f.Scan(&n3)

	//Processamento
	N = n1*100 + n2*10 + n3
	Q = N * N

	//Saída
	if (n1 <= 0) || (n1 > 9) || (n2 < 0) || (n3 > 9) || (n3 < 0) || (n3 > 9) {
		f.Println("Digíto inválido")
	} else {
		f.Println("O número e o seu quadrado são respectivamente: ", N, ",", Q)
	}

}
