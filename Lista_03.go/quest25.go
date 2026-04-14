package main

import (
	f "fmt"
	m "math"
)

func main() {

	var S float64

	f.Println("O valor do somatório S = 1/225 - 2/196 + 4/169 - 8/144 + ... + 16384/1 é: ")

	for i := 1; i <= 15; i++ {
		a := m.Pow(2, float64(i-1))
		b := m.Pow(float64(16-i), 2)

		if i%2 == 0 {
			S -= a / b
		} else {
			S += a / b
		}
	}

	f.Printf("S = %.4f", S)

}
