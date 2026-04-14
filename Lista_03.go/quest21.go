package main

import f "fmt"

func main() {

	var b, n int

	f.Println("Insira dois valores positivos, uma para a base (b) e o outro para o expoente (n), respectivamente: ")
	f.Println("Coloque os valores sendo b >= 2 e n > 1")
	f.Scan(&b, &n)

	if b < 2 || n <= 1 {
		f.Println("Algum dos valores inseridos está incorreto!")
		return
	}

	P := potencia(b, n)

	f.Printf("O número %d elevado a %d resulta em: %d", b, n, P)
}

func potencia(b, n int) int {
	r := 1

	for i := 1; i <= n; i++ {
		r *= b
	}
	return r
}
