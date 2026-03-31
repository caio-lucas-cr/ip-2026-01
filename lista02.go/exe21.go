package main

import f "fmt"

func main() {

	var id, n1, n2, n3, me, mf float32
	var conceito, status string

	f.Println("Insira o seu número de identificação: ")
	f.Scan(&id)

	f.Println("Insira o valor da nota da primeira verificação: ")
	f.Scan(&n1)

	f.Println("Insira o valor da nota da segunda verificação: ")
	f.Scan(&n2)

	f.Println("Insira o valor da nota da terceira verificação: ")
	f.Scan(&n3)

	f.Println("Insira a média do exercícios: ")
	f.Scan(&me)

	mf = (n1 + 2*n2 + 3*n3 + me) / 7

	if mf >= 9.0 && mf <= 10.0 {
		conceito = "A"
	} else if mf >= 7.5 && mf < 9.0 {
		conceito = "B"
	} else if mf >= 6.0 && mf < 7.5 {
		conceito = "C"
	} else if mf >= 4.0 && mf < 6.0 {
		conceito = "D"
	} else if mf < 4 {
		conceito = "E"
	}

	if conceito == "A" || conceito == "B" || conceito == "C" {
		status = "APROVADO"
	} else {
		status = "REPROVADO"
	}

	f.Println("Número de identificação: ", id)
	f.Printf("Nota 1: %.1f", n1)
	f.Println(" ")
	f.Printf("Nota 2: %.1f", n2)
	f.Println(" ")
	f.Printf("Nota 3: %.1f", n3)
	f.Println(" ")
	f.Printf("Média do exercícios: %.1f", me)
	f.Println(" ")
	f.Printf("Média final de aproveitamento: %.1f", mf)
	f.Println(" ")
	f.Println("Conceito: ", conceito)
	f.Println("Status: ", status)

}
