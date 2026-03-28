package main

import f "fmt"

func main() {
	//Algoritmo "Equação do Delta na Equação de Báskara"

	var A, B, C, delta float32

	//Entrada
	f.Println("Digite o coeficiente A: ")
	f.Scan(&A)

	f.Println("Digite o coeficiente B: ")
	f.Scan(&B)

	f.Println("Digite o coeficiente C: ")
	f.Scan(&C)

	//Processamento
	delta = B*B - 4*A*C

	//Saída
	f.Printf("O VALOR DE DELTA E: %.2f ", delta)

}
