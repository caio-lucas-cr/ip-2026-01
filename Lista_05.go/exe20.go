package main

import f "fmt"

func main() {
	var jogadas [20]int
	var frequencia [7]int

	f.Printf("Digite os números sorteados em 20 jogadas (de 1 a 6):\n")

	for i := 0; i < len(jogadas); i++ {
		var valor int
		f.Printf("Jogada %d: ", i+1)
		f.Scan(&valor)

		if valor >= 1 && valor <= 6 {
			jogadas[i] = valor
			frequencia[valor]++ 
		} else {
			f.Println("Valor inválido! Digite um número entre 1 e 6.")
			i-- 
		}
	}

	f.Println("\n--- Números Sorteados ---")
	f.Println(jogadas)

	f.Println("\n--- Frequência das Faces ---")
	for face := 1; face <= 6; face++ {
		vezes := frequencia[face]
		percentual := (float64(vezes) / float64(20)) * 100
		f.Printf("Face %d: %d vezes (%.1f%%)\n", face, vezes, percentual)
	}
}