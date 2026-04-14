package main

import f "fmt"

func main() {
	var ingresso int = 130
	var p, maximo, lucro, melhorPreco float64
	var melhorIngresso int
	primeiro := true

	f.Println("O melhor resultado para a companhia de teatro é: ")

	for p = 6.0; p >= 1.0; p -= 0.6 {
		lucro = (p * float64(ingresso)) - 300

		if primeiro {
			maximo = lucro
			melhorPreco = p
			melhorIngresso = ingresso
			primeiro = false
		} else if lucro > maximo {
			maximo = lucro
			melhorPreco = p
			melhorIngresso = ingresso
		}

		ingresso += 30
	}

	f.Printf("\n--- RESULTADO ---")
	f.Printf("\nO lucro máximo é: R$%.2f", maximo)
	f.Printf("\nO melhor preço por ingresso é: R$%.2f", melhorPreco)
	f.Printf("\nA quantidade de ingressos é: %d\n", melhorIngresso)
}
