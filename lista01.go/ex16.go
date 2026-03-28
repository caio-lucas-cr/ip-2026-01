package main

import f "fmt"

func main() {

	//Algoritmo "Reajuste salarial"

	var salario, reaj_sal float32

	//Entrada
	f.Println("Digite o valor do salário: ")
	f.Scan(&salario)

	//Processamento
	if salario <= 300 {
		reaj_sal = salario * 1.5
	} else {
		reaj_sal = salario * 1.3
	}

	//Saída
	f.Printf("SALARIO COM REAJUSTE: %.2f", reaj_sal)
}
