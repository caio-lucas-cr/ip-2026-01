package main

import f "fmt"

func main() {

	var N1, N2 int

	f.Println("Digite dois números naturais, distintos : ")
	f.Scan(&N1, &N2)

	if (N1 <= 0) || (N2 <= 0) || (N1 == N2) {
		f.Println("O intervalo inserido não é válido!")
		return
	}

	if N1 > N2 {
		N1, N2 = N2, N1
	}

	existe := false

	f.Printf("\nNúmeros primos entre %d e %d: \n", N1, N2)

	for i := N1 + 1; i < N2; i++ {
		primo := true

		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				primo = false
				break
			}
		}

		if primo {
			f.Printf("%d ", i)
			existe = true
		}

	}

	if !existe {
		f.Println("Não existem números primos nesse intervalo!")
	}
}
