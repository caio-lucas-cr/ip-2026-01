package main

import (
	f "fmt"
	m "math"
)

func main() {

	var milho uint

	f.Println("A quantidade de milho que pode-se colocar no tabuleiro de xadrez é dada por: ")

	for i := 0; i <= 63; i++ {
		qtd := m.Pow(2, float64(i))

		milho += uint(qtd)
	}

	f.Printf("\nA quantidade de milho é: %d", milho)
}
