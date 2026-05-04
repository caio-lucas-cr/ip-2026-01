package main

import f "fmt"

func main() {

	vetor := make([]int, 10)

	f.Println("Insira 10 números inteiros distintos:")
	
	for i := 0; i < len(vetor); i++ {
		f.Printf("Posição %d: ", i)
		f.Scan(&vetor[i])
	}

	menorElemento := vetor[0]
	posicao := 0

	for i := 1; i < len(vetor); i++ {
		if vetor[i] < menorElemento {
			menorElemento = vetor[i]
			posicao = i
		}
	}

	f.Printf("\nO menor elemento do vetor é: %d\n", menorElemento)
	f.Printf("Sua posição dentro do vetor é: %d\n", posicao)
}