package main

import (
	"fmt"
	"math"
)

func main() {
	//Algoritmo "Volume da Pirâmide de Base Hexagonal"

	var v, a, h float64

	//Entrada
	fmt.Println("Digite a altura da pirâmide: ")
	fmt.Scan(&h)

	fmt.Println("Digite a aresta do hexágono: ")
	fmt.Scan(&a)

	//Processamento
	v = ((a * a) * (math.Sqrt(3)) * h) / 2

	//Saída
	fmt.Printf("O VOLUME DA PIRÂMIDE E: %.2f METROS CUBICOS", v)

}
