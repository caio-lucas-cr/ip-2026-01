package main

import f "fmt"

func main() {

	//Algoritmo: "Conversão de temperatura"

	var n int
	var fa, cel float64

	//Entrada
	f.Println("Digite o número de temperaturas em Fahrenheit a serem convertidas: ")
	f.Scan(&n)

	//Saída
	if n <= 0 {
		f.Println("Quantidade inválida!")
	}

	for i := 1; i <= n; i++ {
		f.Println("Digite a temperatura", i, ":")
		f.Scan(&fa)

		cel = 5 * (fa - 32) / 9

		f.Printf(" %.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS ", fa, cel)
		f.Println(" ")
		f.Println(" ")
	}

}
