package main

import f "fmt"

func main() {

	var n1, n2, s float64

	f.Println("Digite o 1° número: ")
	f.Scan(&n1)

	f.Println("Digite o 2° número: ")
	f.Scan(&n2)

	s = soma(n1, n2)

	f.Printf("A soma entre %.1f e %.1f é %.1f", n1, n2, s)

}

func soma(n1, n2 float64) float64 {
	return n1 + n2
}
