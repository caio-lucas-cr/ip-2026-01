package main

import f "fmt"

func main() {
	var vetor [10]int

	f.Println("Digite 10 números inteiros (serão ordenados durante a leitura):")

	for i := 0; i < len(vetor); i++ {
		var n int
		f.Printf("Digite o %dº número: ", i+1)
		f.Scan(&n)

		posicao := i

		for posicao > 0 && vetor[posicao-1] > n {
			vetor[posicao] = vetor[posicao-1]
			posicao--
		}

		vetor[posicao] = n
	}

	f.Println("\nVetor final totalmente ordenado:")
	f.Println(vetor)
}
