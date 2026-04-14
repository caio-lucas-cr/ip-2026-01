package main

import f "fmt"

func main() {

	var n, produto, k int

	f.Println("Insira um número inteiro positivo: ")
	f.Scan(&n)

	if n <= 0 {
		f.Println("O valor inserido é inválido!")
		return
	}

	triangular := false

	k = 1

	for {
		produto = k * (k + 1) * (k + 2)

		if produto == n {
			triangular = true
			f.Printf("\nO número %d é um número trinagular!\n", n)
			f.Printf("Sendo que %d é %dx%dx%d", n, k, k+1, k+2)
			break
		}

		if produto > n {
			break
		}

		k++

	}

	if !triangular {
		f.Printf("\nO número %d NÃO é um número trinagular!", n)
	}
}
