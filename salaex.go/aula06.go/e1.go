package main

import f "fmt"

func main() {

	var nota [3]float64
	var soma, media float64

	for i := 0; i < len(nota); i++ {
		f.Printf("Insira a %d° nota: \n", i+1)
		f.Scan(&nota[i])
	}

	for i := 0; i <= 2; i++ {
		soma = soma + nota[i]
	}

	media = soma / 3

	f.Printf("O valor da média geral é: %.2f\n", media)

	for i := 0; i <= 2; i++ {
		if nota[i] > media {
			f.Printf("A %d° nota %.2f é maior que a média %.2f\n", i+1, nota[i], media)
		}
	}

}
