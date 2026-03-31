package main

import (
	f "fmt"
	m "math"
	"math/cmplx"
)

func main() {

	var a, b, c float64

	f.Println("Digite o coeficiente A: ")
	f.Scan(&a)
	f.Println("Digite o coeficiente B: ")
	f.Scan(&b)
	f.Println("Digite o coeficiente C: ")
	f.Scan(&c)

	d := m.Pow(b, 2) - 4*(a*c)

	if a == 0 {
		f.Println("O coeficiente A não pode ser zero")
		return
	}
	if d > 0 {
		x1 := (-b + m.Sqrt(d)) / (2 * a)
		x2 := (-b - m.Sqrt(d)) / (2 * a)

		f.Println("As raízes são reais")
		f.Println("São:", x1, "e", x2)
	}
	if d == 0 {
		x1 := -b / (2 * a)

		f.Println("As raízes são reais")
		f.Println("É:", x1)
	}
	if d < 0 {
		dC := complex(d, 0)
		bC := complex(b, 0)
		aC := complex(a, 0)

		x1 := (-bC + cmplx.Sqrt(dC)) / (2 * aC)
		x2 := (-bC - cmplx.Sqrt(dC)) / (2 * aC)

		f.Println("As raízes são imaginárias")
		f.Printf("x1=%.2f % +.2fi", real(x1), imag(x1))
		f.Println("")
		f.Printf("x2=%.2f % +.2fi", real(x2), imag(x2))

	}

}
