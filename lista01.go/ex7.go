package main

import f "fmt"

func main() {

	//Algoritmo "Conversões para o Sistema Métrico"

	var fa, cel, p, m float32

	//Entrada
	f.Println("Digite a temperatura em Fahrenheit: ")
	f.Scan(&fa)

	f.Println("Digite a quantidade de chuva em polegada: ")
	f.Scan(&p)

	//Preocessamento
	cel = 5 * (fa - 32) / 9
	m = 25.4 * p

	//Saída
	f.Printf("O valor em celsius: %.2f ", cel)
	f.Printf("A quantidade de chuva é: %.2f ", m)

}
