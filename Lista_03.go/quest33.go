package main

import f "fmt"

func main() {

	var N1, N2, resto, quociente int

	f.Println("O quociente e o resto da divisão de N1 por N2 (Inteiros positivos) são dados por: ")
	f.Println("Insira o 1º número (N1): ")
	f.Scan(&N1)

	f.Println("Insira o 2º número (N2): ")
	f.Scan(&N2)

	if N1 < 0 || N2 <= 0 {
		f.Println("Algum número inserido é inválido!")
		return
	}

	origN1 := N1

	for i := 1; i <= N1; i++ {
		resto = N1 - N2
		N1 = resto

		if resto < N2 {
			quociente = i
			break
		}
	}

	f.Printf("Quociente(%d,%d) = %d e o Resto(%d,%d) = %d", origN1, N2, quociente, origN1, N2, resto)

}
