package main

import f "fmt"

func main() {
	var x, n int
	f.Println(" Insira dois valores positivos, uma para a base e o outro para o expoente, respectivamente: ")
	f.Scan(&x, &n)

	P := potencia(x, n)

	f.Println("O resultado é: ", P)

}

func potencia(x, n int) int {

	if n == 0 {
		return 1
	}
	return x * potencia(x, n-1)
}
