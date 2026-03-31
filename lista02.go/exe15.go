package main

import f "fmt"

func main() {

	var d, m, a int

	f.Println("Insira o valor númerico do dia: ")
	f.Scan(&d)

	f.Println("Insira o valor númerico do mês: ")
	f.Scan(&m)

	f.Println("Insira o valor númerico do ano: ")
	f.Scan(&a)

	mes := []string{
		" ", "Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro",
	}

	f.Printf("%d de %s de %d", d, mes[m], a)
}
