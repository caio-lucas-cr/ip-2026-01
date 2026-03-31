package main

import f "fmt"

func main() {

	var n int

	f.Println("Digite um número inteiro: ")
	f.Scan(&n)

	if n == 0 {
		f.Print("O número é nulo")
	} else if n > 0 {
		f.Print("O número é positivo")
	} else if n < 0 {
		f.Print("O número é negativo")
	}
}
