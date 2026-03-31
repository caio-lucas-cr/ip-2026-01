package main

import f "fmt"

func main() {

	var numInt [5]int
	var soma int

	for i := 0; i < len(numInt); i++ {
		f.Printf("Insira o %d° valor inteiro: \n", i+1)
		f.Scan(&numInt[i])
	}

	for i := 0; i < len(numInt); i++ {
		soma += numInt[i]
	}

	f.Println("A soma desses valores é igual a: ", soma)

}