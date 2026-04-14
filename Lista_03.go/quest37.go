package main

import (
	f "fmt"
	m "math"
	sc "strconv"
)

func main() {
	var base8 string
	var base10 int

	f.Print("Digite um número na base 8 para tranformá-lo na base 10: ")
	f.Scan(&base8)

	for i := 0; i < len(base8); i++ {
		caractere := string(base8[i])

		digito, err := sc.Atoi(caractere)

		if err != nil || digito < 0 || digito > 7 {
			f.Println("O número inserido não é um número na base 8 válido.")
			return
		}

		potencia := (len(base8) - 1) - i

		base10 += digito * int(m.Pow(8, float64(potencia)))
	}

	f.Printf("O número %s na base 10 é: %d", base8, base10)
}
