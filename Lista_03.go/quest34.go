package main

import f "fmt"

func main() {

	var N1, N2, MMC int

	f.Println("O mínimo múltiplo comum de N1 e N2 é dado por: ")
	f.Println("Insira o 1º número (N1): ")
	f.Scan(&N1)

	f.Println("Insira o 2º número (N2): ")
	f.Scan(&N2)

	origN1 := N1
	origN2 := N2

	if N1 > N2 {
		N1, N2 = N2, N1
	}

	for i := N1; i <= N1*N2; i++ {
		if i%N1 == 0 && i%N2 == 0 {
			MMC = i
			break
		}
	}

	f.Printf("MMC(%d,%d) = %d", origN1, origN2, MMC)

}
