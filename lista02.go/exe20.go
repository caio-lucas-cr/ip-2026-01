package main

import f "fmt"

func main() {

	var p, v, m float64
	var cond string

	f.Println("Digite o preço de etiqueta do produto: ")
	f.Scan(&p)

	f.Println("Selecione a condição de pagamento do produto: ")
	f.Println("1- À vista, dinheiro ou cheque, 10% de desconto")
	f.Println("2- À vista, cartão de crédito, 5% de desconto")
	f.Println("3- Em 2 vezes, preço normal de etiqueta sem juros")
	f.Println("4- Em 3 vezes, preço normal de etiqueta + 10% de juros")
	f.Scan(&cond)

	if cond == "1" {
		v = 0.9 * p
		f.Printf("O valor a ser pago em dinheiro, à vista, é de R$%.2f", v)
	} else if cond == "2" {
		v = 0.95 * p
		f.Printf("O valor a ser pago no cartão de crédito, à vista, é de R$%.2f", v)
	} else if cond == "3" {
		v = p
		m = p / 2
		f.Printf("O valor a ser pago é de R$%.2f, em duas parcelas de R$%.2f", v, m)
	} else if cond == "4" {
		v = 1.1 * p
		m = 1.1 * p / 3
		f.Printf("O valor a ser pago é de R$%.2f, em três parcelas de R$%.2f", v, m)
	}

}
