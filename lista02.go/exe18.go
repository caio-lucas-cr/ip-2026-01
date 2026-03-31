package main

import f "fmt"

func main() {

	var p, v float32
	var cat, dia string

	f.Println("DIgite o preço normal do DVD: ")
	f.Scan(&p)

	f.Println("Digite o dia da semana do aluguel: ")
	f.Scan(&dia)

	f.Println("Digite a categoria do DVD alugado: ")
	f.Scan(&cat)

	if dia == "segunda" || dia == "terça" || dia == "quinta" {
		if cat == "comum" {
			v = 0.6 * p
			f.Printf("O preço final a ser pago pela locação do DVD é de: %.2f", v)
		} else {
			v = 0.75 * p
			f.Printf("O preço final a ser pago pela locação do DVD é de: %.2f", v)
		}
	} else if dia == "quarta" || dia == "sexta" || dia == "sábado" || dia == "domingo" {
		if cat == "comum" {
			v = p
			f.Printf("O preço final a ser pago pela locação do DVD é de: %.2f", v)
		} else {
			v = 1.15 * p
			f.Printf("O preço final a ser pago pela locação do DVD é de: %.2f", v)
		}
	} else {
		f.Println("Algum dos campos de preenchimento foi inserido com informações equivocadas!")
	}

}
