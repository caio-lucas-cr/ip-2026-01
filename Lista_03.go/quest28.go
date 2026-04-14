package main

import (
	f "fmt"
	m "math"
)

func main() {

	var S, pi float64

	f.Println("Sendo S = 1/1³ - 1/3³ + 1/5³ - 1/7³+ 1/9³ - ..., e ∏ = (S*32)^-3")
	f.Println("O valor de ∏, (tendo S 51 termos) é: ")

	base := -1
	for i := 1; i <= 51; i++ {
		base = base + 2

		denominador := m.Pow(float64(base), 3)

		if i%2 == 0 {
			S -= 1 / denominador
		} else {
			S += 1 / denominador
		}
	}

	pi = m.Pow(S*32, 1.0/3.0)

	f.Printf(" ∏ = %.10f", pi)

}
