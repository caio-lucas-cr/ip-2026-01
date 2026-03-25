package main

import "fmt"

func main() {

	var a, b, c float64

	var valido bool

	fmt.Println("Digite os lados do triângulo: ")
	fmt.Scan(&a, &b, &c)

	if (a+b > c) && (a+c > b) && (b+c > a) {
		fmt.Println("É um triângulo válido!")
		valido = true
	} else {
		fmt.Println("O triângulo não é válido")
		valido = false
	}

	if valido == true {

		if (a != b) && (a != c) && (b != c) {
			fmt.Println("O triângulo é Escaleno")
		}

		if (a == b) && (a != c) || (a == c) && (a != b) || (b == c) && (b != a) {
			fmt.Println("O triângulo é Isósceles")
		}

		if (a == b) && (a == c) {
			fmt.Println("O triângulo é equilátero")
		}
	} else {
		fmt.Println("Os valores dos lados inseridos não formam um triângulo.")
	}

}
