package main

import f "fmt"

func main() {
	idades := make([]int, 50)
	contagem := make(map[int]int)

	f.Printf("Insira as 50 idades:\n")

	for i := 0; i < len(idades); i++ {
		f.Printf("Idade %d: ", i+1)
		f.Scan(&idades[i])
		
		contagem[idades[i]]++
	}

	moda := idades[0]
	maiorFrequencia := 0

	for idade, frequencia := range contagem {
		if frequencia > maiorFrequencia {
			maiorFrequencia = frequencia
			moda = idade
		}
	}

	f.Printf("\n--- Resultado ---")
	f.Printf("\nA moda das idades é: %d", moda)
	f.Printf("\nEssa idade apareceu %d vezes.\n", maiorFrequencia)
}