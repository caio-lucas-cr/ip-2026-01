package main

import f "fmt"

func main() {

	var a, b, Soma, k, z int

	f.Println("Digite dois números ineteiros: ")
	f.Scan(&a, &b)

	Soma = a + b

	if Soma > 20 {
		k = Soma + 8
		f.Println("O valor é :", k)
	} else if Soma <= 20 {
		z = Soma - 5
		f.Println("O valor é ", z)
	}
}