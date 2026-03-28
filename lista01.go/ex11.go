package main

import f "fmt"

func main() {
	//Algoritmo "Divísivel de 3 e 5"

	var N int

	//Entrada
	f.Println("Digite o número inteiro: ")
	f.Scan(&N)

	//Saída
	if (N%3 == 0) && (N%5 == 0) {
		f.Println("O NUMERO E DIVISIVEL")
	} else {
		f.Println("O NUMERO NAO E DIVISIVEL")
	}

}
