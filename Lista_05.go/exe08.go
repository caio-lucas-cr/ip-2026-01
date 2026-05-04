package main

import (
	f "fmt"
	m "math"
)

func main() {

	vetor := make([]int, 15)
	var raiz []float64
	var calculo float64

	f.Println("Insira 15 números inteiros:")

	for i := 0; i < len(vetor); i++ {
		f.Printf("Vetor[%d]: ", i)
		f.Scan(&vetor[i])
	}

	for i := 0; i < len(vetor); i++ {

		if vetor[i] < 0 {
			calculo = -1
		} else {
			calculo = m.Sqrt(float64(vetor[i]))
		}

		raiz = append(raiz, calculo)
	}

	f.Println("\nVetor raíz:")
	f.Printf("%.2f ", raiz)
}
