package main

import f "fmt"

func main() {

	//Algoritmo: "Conta de água"

	var u int
	var cons_mc, v float32
	var categoria string

	//Entrada
	f.Println("Digite a conta do cliente: ")
	f.Scan(&u)

	f.Println("Digite o consumo de água por metros cúbicos: ")
	f.Scan(&cons_mc)

	f.Println("Digite o tipo de consumidor: ")
	f.Scan(&categoria)

	//Processamento
	if categoria == "residencial" {
		v = 5 + 0.05*cons_mc
	} else if categoria == "comercial" {
		if cons_mc <= 80 {
			v = 500
		} else {
			v = 500 + 0.25*(cons_mc-80)
		}
	} else if categoria == "industrial" {
		if cons_mc <= 100 {
			v = 800
		} else {
			v = 800 + 0.04*(cons_mc-100)
		}
	}

	//Saída
	f.Println("Conta: ", u)
	f.Printf("Valor a ser pago: R$ %.2f ", v)

}
