package main

import (
	f "fmt"
	"strconv"
)

func main() {
	var cpf string
	var soma9, soma10, peso1, peso2, digito1_original, digito2_original int

	f.Println("Insira o cpf:")
	f.Scan(&cpf)
	soma9 = 0
	peso1 = 10
	soma10 = 0
	peso2 = 11
	for i := 0; i < 9; i++ {
		digito, _ := strconv.Atoi(string(cpf[i]))
		soma9 = soma9 + digito*peso1
		peso1--
	}
	resto1 := soma9 % 11
	digito1_calculado := 0
	if resto1 < 2 {
		digito1_calculado = 0
	}
	if resto1 >= 2 {
		digito1_calculado = 11 - resto1
	}
	for j := 0; j < 9; j++ {
		digito, _ := strconv.Atoi(string(cpf[j]))
		soma10 = soma10 + digito*peso2
		peso2--
	}
	resto2 := (soma10 + 2*digito1_calculado) % 11
	digito2_calculado := 0

	if resto2 < 2 {
		digito2_calculado = 0
	}
	if resto2 >= 2 {
		digito2_calculado = 11 - resto2
	}
	digito1_original, _ = strconv.Atoi(string(cpf[9]))

	digito2_original, _ = strconv.Atoi(string(cpf[10]))

	if digito1_calculado != digito1_original || digito2_calculado != digito2_original {
		f.Println("O CPF digitado não é válido")
	} else {
		f.Println("O CPF é válido")
	}
}
