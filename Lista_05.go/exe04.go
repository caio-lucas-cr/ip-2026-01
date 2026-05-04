package main

import f "fmt"

func main() {

	A := make([]int, 10)
	contagem := make(map[int]int)

	f.Println("Digite 10 números inteiros para o vetor A:")

	for i := 0; i < len(A); i++ {
		f.Printf("A[%d]: ", i)
		f.Scan(&A[i])

		contagem[A[i]]++
	}

	f.Println("\n--- Elementos Repetidos ---")

	encontrouRepetido := false

	for numero, qtd := range contagem {
		if qtd > 1 {
			f.Printf("O número %d se repete %d vezes.\n", numero, qtd)
			encontrouRepetido = true
		}
	}

	if !encontrouRepetido {
		f.Println("Não foram encontrados números repetidos.")
	}
}