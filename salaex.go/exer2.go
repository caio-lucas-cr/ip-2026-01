package main

import "fmt"

func main() {

	var n1, n2, p1, p2, media float64

	fmt.Println("Digite as notas n1 e n2, respectivamente: ")
	fmt.Scan(&n1, &n2)

	fmt.Println("Digite os pesos p1 e p2, respectivamente: ")
	fmt.Scan(&p1, &p2)

	media = ((n1 * p1) + (n2 * p2)) / (p1 + p2)

	fmt.Println("A média das notas com seus respectivos pesos é: ", media)

}
