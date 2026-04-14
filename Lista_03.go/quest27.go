package main

import (
	f "fmt"
	m "math"
)

func main() {

	var x, cos, k float64

	f.Println("Sendo cosseno(x) = 1 - x^2/2! + x^4/4! - x^6/6! + x^8/8! - ...")
	f.Println("Insira o valor de x para calcular o resultado considerando 20 elementos dessa série: ")
	f.Scan(&x)

	cos = 1

	for i := 1; i <= 19; i++ {
		expoente := i * 2

		k = 1

		for j := 1; j <= expoente; j++ {
			k = k * float64(j)
		}

		y := m.Pow(x, float64(expoente))

		if i%2 == 0 {
			cos += y / float64(k)
		} else {
			cos -= y / float64(k)
		}

	}

	f.Printf("\ncosseno(%.1f) = %.6f", x, cos)

}
