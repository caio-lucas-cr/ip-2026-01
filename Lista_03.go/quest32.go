package main

import f "fmt"

func main() {

	var N1, N2, produto int

	f.Println("Digite o 1º número (N1): ")
	f.Scan(&N1)

	f.Println("Digite o 2º número (N2): ")
	f.Scan(&N2)

	origN1, origN2 := N1, N2

	if (N1 > 0) && (N2 < 0) {
		N1, N2 = N2, N1
	}

	if (N1 < 0) && (N2 < 0) {
		N1, N2 = -N1, -N2
	}

	for i := 0; i < N2; i++ {
		produto += N1
	}

	f.Printf("Assim: %d x %d = %d", origN1, origN2, produto)

}
