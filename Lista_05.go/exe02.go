package main

import f "fmt"

func main() {

	vetor1 := make([]int, 10)
	vetor2 := make([]int, 5)
	var somaVetor2 int

	f.Println("Digite os 10 números do primeiro vetor:")
	for i := 0; i < len(vetor1); i++ {
		f.Printf("Vetor 1 [%d]: ", i)
		f.Scan(&vetor1[i])
	}

	f.Println("\nDigite os 5 números do segundo vetor:")
	for i := 0; i < len(vetor2); i++ {
		f.Printf("Vetor 2 [%d]: ", i)
		f.Scan(&vetor2[i])
		somaVetor2 += vetor2[i]
	}

	var resultantePares []int
	var resultanteImpares []int

	for i := 0; i < len(vetor1); i++ {
		num := vetor1[i] 
		if num%2 == 0 {
			resultantePares = append(resultantePares, num+somaVetor2)
		} else {
			resultanteImpares = append(resultanteImpares, num+somaVetor2)
		}
	}

	f.Println("\n--- RESULTADOS ---")
	f.Printf("Vetor 1: %v\n", vetor1)
	f.Printf("Vetor 2: %v | Soma: %d\n", vetor2, somaVetor2)
	
	f.Printf("\nPrimeiro vetor resultante (Pares): %v\n", resultantePares)
	f.Printf("Segundo vetor resultante (Ímpares): %v\n", resultanteImpares)
}