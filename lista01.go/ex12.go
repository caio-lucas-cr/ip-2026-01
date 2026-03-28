package main

import f "fmt"

func main() {
	//Algoritmo "Locadora de charretes"

	var h int
	var x int

	//Entrada
	f.Println("Digite a quantidade de horas que utilizou a charrete: ")
	f.Scan(&h)

	//Processamento
	x = (h/3)*10 + (h%3)*5

	//Saída
	f.Println("O VALOR A PAGAR E: ", x)

}
