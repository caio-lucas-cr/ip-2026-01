package main

import f "fmt"

func main() {

	var n, x, y, z int

	f.Println("Digite a quantidade de termos (N >= 3): ")
	f.Scan(&n)

	if n < 3 {
		f.Println("Quantidade de termos inválida!")
		return
	}

	x = 0
	y = 1

	f.Printf("Sequência de Fibonacci: %d - %d", x, y)

	for i := 3; i <= n; i++ {
		z = x + y
		f.Printf(" - %d", z)

		x, y = y, z

	}
}
