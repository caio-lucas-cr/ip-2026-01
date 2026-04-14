package main

import f "fmt"

func main() {

	var sal_c, sal_j float64

	f.Println("Insira o salário inicial de Carlos: ")
	f.Scan(&sal_c)

	sal_j = sal_c / 3

	meses := 0

	for sal_j < sal_c {
		sal_c += sal_c * 0.02
		sal_j += sal_j * 0.05
		meses++
	}

	f.Printf("Serão necessários %d meses para o salário de João igualar ou ultrapassar o salário de Carlos.\n", meses)
	f.Printf("O salário final de João é: R$%.2f\n", sal_j)
	f.Printf("O salário final de Carlos é: R$%.2f\n", sal_c)
}
