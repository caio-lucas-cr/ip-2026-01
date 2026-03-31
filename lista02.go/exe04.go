package main

import (
	"fmt"
	"math"
)

func main() {

	var n, q, r float64

	fmt.Println("Digite um número: ")
	fmt.Scan(&n)

	r = math.Sqrt(n)
	q = math.Pow(n, 2)

	if n >= 0 {
		fmt.Printf("O valor é: %.2f", r)
	} else if n < 0 {
		fmt.Println("O valor é: ", q)
	}

}