package main

import (
	f "fmt"
	m "math"
)

func main() {

	var SenA float64

	f.Printf("%-6s | %s\n", "Ângulo (A)", "Seno(A)")
	f.Println("----------------------------")

	for A := 0.0; A <= 6.3; A += 0.1 {

		A3 := m.Pow(A, 3)
		A5 := m.Pow(A, 5)
		A7 := m.Pow(A, 7)

		SenA = A - A3/6 + A5/120 - A7/5040

		f.Printf("A = %-6.1f | sen(%.1f) = %.6f\n", A, A, SenA)
	}

}
