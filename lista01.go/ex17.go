package main

import f "fmt"

func main() {

	//Algoritmo "Série de pares"

	var x, y int

	//Entrada
	f.Println(" Digite dois números inteiros: ")
	f.Scan(&x, &y)

	f.Println("--------------------------------")

	//Saída
	if x%2 == 0 {
		for i := 0; i < y; i++ {
			f.Println(x)
			x = x + 2
		}
	} else {
		f.Println("O PRIMEIRO NÚMERO NÃO É PAR")
	}

}
