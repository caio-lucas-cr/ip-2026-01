package main

import f "fmt"

func main() {

	var y, x float64

	f.Println("Digite um número: ")
	f.Scan(&x)

	y = 8/(2 - x)

	f.Printf("f(x) = %.2f",y)
}