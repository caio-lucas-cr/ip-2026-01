package main

import f "fmt"

func main() {

	var numeros [10]int
	encontrou := false

	f.Println("Digite 10 números inteiros:")

	for i := 0; i < 10; i++ {
		f.Printf("Posição [%d]: ", i)
		f.Scan(&numeros[i])
	}

	f.Println("\n--- Resultados ---")

	for i := 0; i < len(numeros); i++ {
		if numeros[i] > 50 {
			f.Printf("Número: %d | Posição: %d\n", numeros[i], i)
			encontrou = true
		}
	}

	if !encontrou {
		f.Println("Não existe nenhum número superior a 50 no vetor.")
	}
}