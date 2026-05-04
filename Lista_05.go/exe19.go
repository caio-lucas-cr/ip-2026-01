package main

import f "fmt"

func main() {
	var numeros [10]int
	var divisores [5]int

	f.Println("Digite os 10 números do primeiro vetor:")

	for i := 0; i < len(numeros); i++ {
		f.Printf("Número [%d]: ", i)
		f.Scan(&numeros[i])
	}

	f.Println("\nDigite os 5 números do vetor de divisores:")
	
	for i := 0; i < len(divisores); i++ {
		f.Printf("Divisor [%d]: ", i)
		f.Scan(&divisores[i])
	}

	f.Println("\n--- Resultado da Análise ---")

	for i := 0; i < len(numeros); i++ {
		f.Printf("\nNúmero %d:", numeros[i])
		encontrouDivisor := false

		for j := 0; j < len(divisores); j++ {
			if divisores[j] == 0 {
				continue
			}

			if numeros[i]%divisores[j] == 0 {
				f.Printf("\n  Divisível por %d na posição %d", divisores[j], j)
				encontrouDivisor = true
			}
		}

		if !encontrouDivisor {
			f.Print("\n  Não possui divisores no segundo vetor.")
		}
		f.Println()
	}
}