package main

import f "fmt"

func main() {
	//Algoritmo "Conversão de Nota em Conceito"

	var nota float32

	//Entrada
	f.Println("Digite a nota: ")
	f.Scan(&nota)

	//Processamento e Saída
	if (nota >= 9.0) && (nota <= 10.0) {
		f.Println("NOTA = ", nota, " ", "CONCEITO = A")
	} else if (nota >= 7.5) && (nota < 9.0) {
		f.Println("NOTA = ", nota, " ", "CONCEITO = B")
	} else if (nota >= 6.0) && (nota < 7.5) {
		f.Println("NOTA = ", nota, " ", "CONCEITO = C")
	} else {
		f.Println("NOTA = ", nota, " ", "CONCEITO = D")
	}

}
