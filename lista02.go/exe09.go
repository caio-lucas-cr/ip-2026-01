package main

import f "fmt"

func main() {

	var c, v float32

	f.Println(" Digite o valor de compra: ")
	f.Scan(&c)

	if c > 0 && c < 10 {
		v = 1.7 * c
		f.Printf("O valor de venda é: %.2f", v)
	} else if c >= 10 && c < 30 {
		v = 1.5 * c
		f.Printf("O valor de venda é: %.2f", v)
	} else if c >= 30 && c < 50 {
		v = 1.4 * c
		f.Printf("O valor de venda é: %.2f", v)
	} else if c >= 50 {
		v = 1.3 * c
		f.Printf("O valor de venda é: %.2f", v)
	} else if c <= 0 {
		f.Println("O valor de compra é inválido!!")
	}

}
