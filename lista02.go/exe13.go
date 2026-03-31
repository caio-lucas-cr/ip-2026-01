package main

import f "fmt"

func main() {
	
	var n, a, b int

	f.Println("Insira um número inteiro positivo de 3 casas: ")
	f.Scan(&n)

	if n >= 100 && n < 1000 {
		a = n / 100
		b = (n - a*100) / 10
		f.Println("O algarismo da dezena é: ", b)
	} else {
		f.Println("O número não atende as condições exigidas!")
	}
}
