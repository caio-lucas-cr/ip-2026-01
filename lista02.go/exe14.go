package main

import (
	f "fmt"
	s "strings"
)

func main() {

	var v float32
	var resposta string

	const (
		v_ar      = 1750.00
		v_pintura = 800.00
		v_vidro   = 1200.00
		v_direçao = 2000.00
	)

	f.Println("Digite o preço inicial de fábrica do veículo: ")
	f.Scan(&v)

	f.Println("Diga sim ou não para adicionar as opções abaixo: ")

	f.Printf("Ar condicionado (R$%.2f)? (sim/não): ", v_ar)
	f.Scan(&resposta)
	if s.ToLower(resposta) == "sim" {
		v = v + v_ar
	}

	f.Printf("Pintura metálica (R$%.2f)? (sim/não): ", v_pintura)
	f.Scan(&resposta)
	if s.ToLower(resposta) == "sim" {
		v = v + v_pintura
	}

	f.Printf("Vidro elétrico (R$%.2f)? (sim/não): ", v_vidro)
	f.Scan(&resposta)
	if s.ToLower(resposta) == "sim" {
		v = v + v_vidro
	}

	f.Printf("Ar condicionado (R$%.2f)? (sim/não): ", v_direçao)
	f.Scan(&resposta)
	if s.ToLower(resposta) == "sim" {
		v = v + v_direçao
	}

	f.Printf("O valor final do carro é de: %.2f", v)
}
