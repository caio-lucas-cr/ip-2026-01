package main

import "fmt"

func main() {

	var qtd int
	var media, soma, nota float64

	fmt.Println("Insira a quantidade de notas:")
	fmt.Scan(&qtd)

	if qtd <= 0 {
		fmt.Println("Quantidade inválida")
	}

	for i := 1; i <= qtd; i++ {
		fmt.Printf("Digite a nota %d: ", i)
		fmt.Scan(&nota)
		soma += nota
	}
	media = soma / float64(qtd)
	fmt.Println("A média é: ", media)
}
