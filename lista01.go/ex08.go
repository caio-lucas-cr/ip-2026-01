package main

import f "fmt"

func main() {

	//Algoritmo "Custo da Lata de Cerveja"

	var r, a, custo_lata float32

	//Entrada
	f.Println("Digite o valor do raio: ")
	f.Scan(&r)

	f.Println("DIgite o valor da altura: ")
	f.Scan(&a)

	//Processamento
	custo_lata = 100 * (2*(3.14159*(r*r)) + 2*3.14159*r*a)

	//Saída
	f.Printf("O VALOR DO CUSTO E = %.2f ", custo_lata)

}
