package main

import f "fmt"

func main() {

	var x, y, z float64

	f.Println("Digite os três valores para realizar a média: ")
	f.Scan(&x, &y, &z)

	m := media(x, y, z)

	f.Printf("A média dos valores é: %.3f", m)

}

func media(x, y, z float64) float64 {
	return (x + y + z) / 3
}
