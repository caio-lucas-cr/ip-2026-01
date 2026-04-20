package main

import f "fmt"

func main() {

	var reais = []float64{1.2, 3.7, 7.9, 4.6, 5.1}

	som := soma(reais, len(reais)-1)

	f.Printf("A array %v\n", reais)
	f.Printf("A soma dos termos da array é dada por: %.1f", som)

}

func soma(reais []float64, n int) float64 {

	if n == 0 {
		return reais[0]
	}

	return reais[n] + soma(reais, n-1)

}
