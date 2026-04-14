package main

import f "fmt"

func main() {

	var N, a, b, c int

	f.Println("Digite a quantidade de termos (N >= 3): ")
	f.Scan(&N)

	if N < 3 {
		f.Println("Quantidade de termos inválida!")
		return
	}

	f.Println("Digite o 1º termo inteiro: ")
	f.Scan(&a)

	f.Println("Digite o 2º termo inteiro: ")
	f.Scan(&b)

	f.Printf("%d, %d", a, b)

	for i := 3; i <= N; i++ {
		if i%2 == 0 {
			c = b - a
		} else {
			c = b + a
		}
		f.Printf(", %d", c)

		a, b = b, c
	}
}
