package main

import (
	f "fmt"
	m "math"
)

func ehPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	limite := int(m.Sqrt(float64(n)))
	for i := 3; i <= limite; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var vetor [10]int

	f.Println("Digite 10 números inteiros:")

	for i := 0; i < len(vetor); i++ {
		f.Printf("Posição [%d]: ", i)
		f.Scan(&vetor[i])
	}

	f.Println("\n--- Números Primos Encontrados ---")
	encontrou := false

	for i := 0; i < len(vetor); i++ {
		if ehPrimo(vetor[i]) {
			f.Printf("Valor: %d | Posição no vetor: %d\n", vetor[i], i)
			encontrou = true
		}
	}

	if !encontrou {
		f.Println("Nenhum número primo foi digitado.")
	}
}