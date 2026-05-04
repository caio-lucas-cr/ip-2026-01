package main

import f "fmt"

func main() {

	vetor := make([]int, 100)

	for i := 0; i < len(vetor); i++ {
		vetor[i] = 2*i + 1
	}

	f.Printf("Vetor de números ímpares: %v", vetor)
}
