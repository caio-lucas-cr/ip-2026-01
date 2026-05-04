package main

import f "fmt"

func main() {

	altura := make([]float64, 10)
	var somaAlturas, media float64

	f.Println("Insira a altura dos 10 atletas: ")

	for i := 0; i < len(altura); i++ {
		f.Printf("Altura[%d]: ", i)
		f.Scan(&altura[i])
		somaAlturas += altura[i]
	}

	media = somaAlturas / float64(len(altura))

	f.Printf("\n----Resultados----\n")

	for i := 0; i < len(altura); i++ {
		if altura[i] > media {
			f.Printf("A altura[%d]= %.2f é maior que a média %.2f\n", i, altura[i], media)
		}
	}

}
