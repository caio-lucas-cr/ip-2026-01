package main

import (
	f "fmt"
	m "math"
)

func main() {

	var opçao string
	var v, h, r, a float64

	f.Println("Selecione uma das opções que deseja calcular o volume e a área: ")
	f.Println("1-cone / 2-cilindro / 3- esfera")
	f.Scan(&opçao)

	if opçao == "1" {
		f.Println("Digite o raio e a altura do cone, respectivamente: ")
		f.Scan(&r, &h)
		v = m.Pi * r * r * h / 3
		a = m.Pi * r * (m.Sqrt(r*r + h*h))
		f.Printf("O valor do volume do cone é %.2f, e o valor da área é %.2f", v, a)
	} else if opçao == "2" {
		f.Println("Digite o raio e a altura do cilindro, respectivamente: ")
		f.Scan(&r, &h)
		v = m.Pi * r * r * h
		a = 2 * m.Pi * r * h
		f.Printf("O valor do volume do cilindro é %.2f, e o valor da área é %.2f", v, a)
	} else if opçao == "3" {
		f.Println("Digite o raio da esfera: ")
		f.Scan(&r)
		v = 4 * m.Pi * r * r * r / 3
		a = 4 * m.Pi * r * r
		f.Printf("O valor do volume da esfera é %.2f, e o valor da área é %.2f", v, a)
	}

}
