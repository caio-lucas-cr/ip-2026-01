package main

import f "fmt"

func main() {

	var x, y int

	f.Println(" Insira dois valores positivos, uma para a base e o outro para o expoente, respectivamente: ")
	f.Scan(&x, &y)

	P := Potencia(x, y)

	f.Printf("O número %d elevado a %d resulta em: %d", x, y, P)
}

func Potencia(x, y int) int {
	r := 1

	for i := 1; i <= y; i++ {
		r *= x
	}
	return r
}
