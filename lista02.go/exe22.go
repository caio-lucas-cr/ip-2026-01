package main

import f "fmt"

func main() {

	var matricula int
	var qt_he, s_b, s_l float32

	const (
		s_min = 788.00
		v_he  = 10.00
	)

	f.Println("Insira a matrícula do funcionário: ")
	f.Scan(&matricula)

	f.Println("Insira a qunatidade de horas-extras: ")
	f.Scan(&qt_he)

	s_b = 3*s_min + qt_he*v_he

	//De acordo com o enunciado, só o salário mínimo multipllicado por 3 já deixa o salário bruto maior que R$2000,00
	//O desconto de 32% vem de 12% do INSS mais 20% do Imposto de Renda
	if s_b > 2000 {
		s_l = s_b * 0.68
	}

	f.Println("Matrícula do funcionário: ", matricula)
	f.Printf("Salário Bruto: R$%.2f", s_b)
	f.Println(" ")
	f.Printf("Salário Líquido: R$%.2f", s_l)
	f.Println(" ")

}
