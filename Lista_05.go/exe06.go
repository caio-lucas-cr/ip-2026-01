package main

import f "fmt"

func main() {
	vetor := make([]int, 100)

	for i := 0; i < len(vetor); i++ {
		vetor[i] = 100 - i
	}

	f.Println("Elementos do vetor (100 a 1):")

	for i := 0; i < len(vetor); i++ {
		f.Printf("%d ", vetor[i])
	}
}
