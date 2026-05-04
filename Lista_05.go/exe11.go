package main 

import (
	f "fmt"
	m "math"
)

func main() {

	var somatorio float64
	numeros:= make([]float64,100)

	f.Println("Insira 100 números a seguir: ")

	for i:=0; i<len(numeros); i++{
		f.Printf("Número[%d]: ",i)
		f.Scan(&numeros[i])
	}

	for i:=0; i<len(numeros)/2; i++{
		somatorio += m.Pow(numeros[i]-numeros[99-i],3)
	}

	f.Printf("\nSomatório: %.2f", somatorio)
}