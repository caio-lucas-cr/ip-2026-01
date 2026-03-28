package main

import f "fmt"

func main() {

	//Algoritmo "Cálculo do Determinante de uma Matriz Quadrada de Duas Dimensões"

	var a, b, c, d, Det float32

	//Entrada
	f.Println("Digite o valor de a: ")
	f.Scan(&a)

	f.Println("Digite o valor de b: ")
	f.Scan(&b)

	f.Println("Digite o valor de c: ")
	f.Scan(&c)

	f.Println("Digite o valor de d: ")
	f.Scan(&d)

	//Processamento
	Det = (a * d) - (c * b)

	//Saída
	f.Printf("O VALOR DO DETERMINANTE E: %.2f ", Det)

}
