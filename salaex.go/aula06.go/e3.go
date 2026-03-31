package main

import f "fmt"

func main() {

	var valReal [10]float64
	var valInv [10]float64

	for i := 0; i < len(valReal); i++ {
		f.Printf("Insira o %d° valor real: \n", i+1)
		f.Scan(&valReal[i])
	}

	for i := 0; i < len(valReal); i++ {
		valInv[i] = valReal[9-i]
	}

	f.Print("A ordem inversa é: ", valInv)

}
