package main

import (
	f "fmt"
	m "math"
)

func main() {

	var V float64

	f.Println("O raio e o volume das respectivas esferas é:")
	f.Printf("\n%s | %s\n", "Raio (cm)", "Volume (cm³)")
	f.Println("----------------------------")

	for R := 0.0; R <= 20.0; R += 0.5 {
		pi := m.Pi
		R3 := m.Pow(R, 3)

		V = 4.0 / 3.0 * pi * R3

		f.Printf("R = %.1fcm | Volume = %.2fcm³\n", R, V)
	}

}
